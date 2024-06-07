package indexer

import (
	// "context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ArkeoNetwork/airdrop/pkg/db"
	"github.com/ArkeoNetwork/airdrop/pkg/types"
	"github.com/ArkeoNetwork/airdrop/pkg/utils"
	arkutils "github.com/ArkeoNetwork/common/utils"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/rpc/client/http"
	// coretypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

type CosmosIndexerParams struct {
	DB    arkutils.DBConfig
	Chain string
}

type CosmosIndexer struct {
	db          *db.AirdropDB
	tm          *http.HTTP
	lcd         *resty.Client
	chain       *types.Chain
	startHeight int64
	endHeight   int64
}

// a page of delegator/validator pairs with amounts for starting balances
type DelegationPage struct {
	DelegationResponses []DelegationResponse `json:"delegation_responses"`
}

type Delegation struct {
	DelegatorAddress string `json:"delegator_address"`
	ValidatorAddress string `json:"validator_address"`
	Shares           string `json:"shares"`
}

type TransactionAmount struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type TransactionMessages struct {
	Type             string            `json:"@type"`
	DelegatorAddress string            `json:"delegator_address"`
	ValidatorAddress string            `json:"validator_address"`
	Amount           TransactionAmount `json:"amount"`
}

type TransactionBody struct {
	Messages []TransactionMessages `json:"messages"`
	Memo     string                `json:"memo"`
}

type Transaction struct {
	Body TransactionBody `json:"body"`
}

type TransactionsResponse struct {
	Txs []*Transaction `json:"txs"`
	// TxResponse string      `json:"tx_responses"`
	Pagination string `json:"pagination"`
	Total      string `json:"total"`
}

type DelegationResponse struct {
	Delegation Delegation `json:"delegation"`
	Balance    struct {
		Denom  string `json:"denom"`
		Amount int64  `json:"amount,string"`
	} `json:"balance"`
}

type ImportedDelegation struct {
	AppState struct {
		Staking struct {
			Delegations []Delegation `json:"delegations"`
		} `json:"staking"`
	} `json:"app_state"`
}

func NewCosmosIndexer(params CosmosIndexerParams) (*CosmosIndexer, error) {
	d, err := db.New(params.DB)
	if err != nil {
		return nil, errors.Wrapf(err, "error connecting to db")
	}
	chain, err := d.FindChain(params.Chain)
	if err != nil {
		return nil, errors.Wrapf(err, "error finding chain %s", params.Chain)
	}

	log.Infof("connecting to tendermint node at %s", chain.RpcUrl)
	tm, err := arkutils.NewTendermintClient(chain.RpcUrl)
	if err != nil {
		return nil, errors.Wrapf(err, "error creating tendermint client with rpc %s", chain.RpcUrl)
	}
	lcd := resty.New().SetTimeout(10 * time.Second).SetBaseURL(chain.LcdUrl)

	return &CosmosIndexer{db: d, tm: tm, lcd: lcd, chain: chain, startHeight: int64(chain.SnapshotStartBlock), endHeight: int64(chain.SnapshotEndBlock)}, nil
}

func (c *CosmosIndexer) IndexDelegationsFromStateExport(dataDir, chain string, height int64) error {
	stateExportFile := fmt.Sprintf("%s/state-export_%d.json", dataDir, height)
	start := time.Now()
	raw, err := os.ReadFile(stateExportFile)
	if err != nil {
		return errors.Wrapf(err, "error reading file %s", stateExportFile)
	}
	log.Infof("read file %s in %.3f seconds", stateExportFile, time.Since(start).Seconds())
	start = time.Now()
	imported := ImportedDelegation{}
	if err = json.Unmarshal(raw, &imported); err != nil {
		return errors.Wrapf(err, "error unmarshalling file %s", stateExportFile)
	}
	log.Infof("unmarshalled delegations %s in %.3f seconds", stateExportFile, time.Since(start).Seconds())

	events := make([]*types.CosmosStakingEvent, 0, len(imported.AppState.Staking.Delegations))

	start = time.Now()
	for _, d := range imported.AppState.Staking.Delegations {
		value, err := parseShares(d.Shares, c.chain.Decimals)
		if err != nil {
			return errors.Wrapf(err, "%s delegation to %s error parsing shares %s", d.DelegatorAddress, d.ValidatorAddress, d.Shares)
		}
		if value <= 0 {
			log.Warnf("%s delegation to %s with value %f. string shares: %s", d.DelegatorAddress, d.ValidatorAddress, value, d.Shares)
		}
		event := &types.CosmosStakingEvent{
			Chain:       c.chain.Name,
			EventType:   "initial",
			Delegator:   d.DelegatorAddress,
			Validator:   d.ValidatorAddress,
			Value:       value,
			BlockNumber: uint64(height),
			TxHash:      "00000000000000000000000000000000",
			EventIndex:  0,
		}
		events = append(events, event)
	}
	log.Infof("created %d staking events in %.3f seconds", len(events), time.Since(start).Seconds())
	start = time.Now()
	if err = c.db.InsertStakingEvents(events); err != nil {
		return errors.Wrapf(err, "error inserting staking events")
	}
	log.Infof("inserted %d staking events in %.3f seconds", len(events), time.Since(start).Seconds())
	return nil
}

func parseShares(s string, decimals uint8) (float64, error) {
	if !strings.Contains(s, ".") {
		return -1, fmt.Errorf("shares %s does not contain a decimal", s)
	}
	parts := strings.Split(s, ".")
	if len(parts) != 2 {
		return -1, fmt.Errorf("shares %s has more than one decimal", s)
	}
	ishares, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return -1, errors.Wrapf(err, "error parsing shares %d", ishares)
	}

	return utils.FromBaseUnits(ishares, decimals), nil
}

// reads a collection of pageN.json files from the dataDir and inserts them into the db
func (c *CosmosIndexer) IndexStartingDelegations(dataDir string) error {
	dir, err := os.ReadDir(dataDir)
	if err != nil {
		return errors.Wrapf(err, "error reading dir %s", dataDir)
	}
	for _, f := range dir {
		if f.IsDir() {
			log.Infof("%s is a directory, skipping", f.Name())
		}
		if !strings.HasSuffix(f.Name(), ".json") {
			log.Infof("%s is not a json file, skipping", f.Name())
			continue
		}
		if !strings.HasPrefix(f.Name(), "page") {
			log.Infof("%s is not a page file, skipping", f.Name())
			continue
		}

		log.Infof("reading file %s", f.Name())
		raw, err := os.ReadFile(fmt.Sprintf("%s/%s", dataDir, f.Name()))
		if err != nil {
			return errors.Wrapf(err, "error reading file %s", f.Name())
		}

		page := DelegationPage{}
		if err = json.Unmarshal(raw, &page); err != nil {
			return errors.Wrapf(err, "error unmarshalling file %s", f.Name())
		}
		log.Debug(page)

		events := make([]*types.CosmosStakingEvent, 0, len(page.DelegationResponses))

		for _, d := range page.DelegationResponses {
			value := utils.FromBaseUnits(d.Balance.Amount, c.chain.Decimals)
			event := &types.CosmosStakingEvent{
				Chain:       c.chain.Name,
				EventType:   "initial",
				Delegator:   d.Delegation.DelegatorAddress,
				Validator:   d.Delegation.ValidatorAddress,
				Value:       value,
				BlockNumber: c.chain.SnapshotStartBlock - 1,
				TxHash:      "00000000000000000000000000000000",
				EventIndex:  0,
			}
			events = append(events, event)
		}
		if err = c.db.InsertStakingEvents(events); err != nil {
			return errors.Wrapf(err, "error inserting staking events for page %s", f.Name())
		}
	}

	return nil
}

func (c *CosmosIndexer) IndexCosmosDelegators() error {
	startHeight := int64(c.chain.SnapshotStartBlock)
	endHeight := int64(c.chain.SnapshotEndBlock)

	latest, err := c.db.FindLatestIndexedCosmosStakingBlock(c.chain.Name)
	if err != nil {
		return errors.Wrapf(err, "error finding latest indexed block")
	}
	if latest > startHeight {
		log.Infof("found latest indexed block %d, starting at %d", latest, latest-1)
		startHeight = latest - 1
	}

	for i := startHeight; i <= endHeight; i++ {
		if err := c.indexCosmosDelegations(i); err != nil {
			log.Errorf("error indexing delegations at height %d: %+v", i, err)
		}
	}
	return nil
}

func shouldStoreTx(tx tmtypes.Tx, txResults *abcitypes.ResponseDeliverTx) bool {
	for _, evt := range txResults.Events {
		switch evt.GetType() {
		case "delegate":
			return true
		case "unbond":
			return true
		case "redelegate":
			return true
		default:
		}
	}
	return false
}

func (c *CosmosIndexer) handleStakingTx(height int64, tx tmtypes.Tx, txResult *abcitypes.ResponseDeliverTx) error {
	log := log.WithField("height", fmt.Sprintf("%d", height))
	txHash := hashTx(tx)
	evtsSequenced := make([]abcitypes.Event, len(txResult.GetEvents()))
	evtsSeq := int64(0)
	evtsIndexMap := make(map[int64]int64, 1024)
	for i, evt := range txResult.GetEvents() {
		switch evt.GetType() {
		case "delegate":
			evtsSequenced[evtsSeq] = evt
			evtsSeq++
		case "redelegate":
			evtsSequenced[evtsSeq] = evt
			evtsSeq++
		case "unbond":
			evtsSequenced[evtsSeq] = evt
			evtsSeq++
		case "message":
			m := make(map[string]string, len(evt.GetAttributes()))
			for _, attr := range evt.GetAttributes() {
				m[string(attr.GetKey())] = string(attr.GetValue())
			}
			if module, ok := m["module"]; ok && module == "staking" {
				if delegator, ok := m["sender"]; ok {
					log.Debugf("adding delegate event delegator %s", delegator)
					evtsSequenced[evtsSeq] = evt
					evtsIndexMap[evtsSeq] = int64(i)
					evtsSeq++
				}
			}
		}
	}
	stakingEvents := make([]*types.CosmosStakingEvent, 0, len(evtsSequenced)/2)
	evtsSequenced = evtsSequenced[:evtsSeq]
	var stakingEvt *stakingEventWrapper
	for i, evt := range evtsSequenced {
		log.Printf("EVT: %s", evt.GetType())
		m := attributesToMap(evt.GetAttributes())
		if i%2 == 0 {
			// staking event itself: delegate,unbond,redelegate
			_, amount, err := parseAmount(m["amount"], c.chain.Decimals)
			if err != nil {
				log.Errorf("error parsing amount %s: %+v", m["amount"], err)
			}
			var srcValidator, destValidator string
			switch evt.GetType() {
			case "delegate":
				destValidator = m["validator"]
			case "redelegate":
				// FIXEME: need to split in to unbond and delegate events
				destValidator = m["destination_validator"]
				srcValidator = m["source_validator"]
			case "unbond":
				srcValidator = m["validator"]
				amount = -amount
			}

			validator := destValidator
			if evt.GetType() == "unbond" {
				validator = srcValidator
			}

			stakingEvt = &stakingEventWrapper{
				CosmosStakingEvent: types.CosmosStakingEvent{
					EventType:   evt.GetType(),
					Validator:   validator,
					Chain:       c.chain.Name,
					Value:       amount,
					BlockNumber: uint64(height),
					TxHash:      txHash,
				},
				srcValidator: srcValidator,
			}
		} else {
			// should be message+staking module+spender (delegator address)?
			log.Debug(evt)
			if stakingEvt == nil {
				return fmt.Errorf("stakingEvt null, event sequencing (programming) issue")
			}
			evtIndex, ok := evtsIndexMap[int64(i)]
			if !ok {
				log.Warnf("no event index found for %d", i)
			}
			stakingEvt.Delegator = m["sender"]
			stakingEvt.EventIndex = evtIndex
			stakingEvents = append(stakingEvents, &stakingEvt.CosmosStakingEvent)
			if stakingEvt.EventType == "redelegate" {
				unbondEvt := types.CosmosStakingEvent{
					EventType:   stakingEvt.EventType,
					EventIndex:  0, // using zero here for extra row to maintain uniqueness
					Delegator:   stakingEvt.Delegator,
					Validator:   stakingEvt.srcValidator,
					Chain:       c.chain.Name,
					Value:       -stakingEvt.Value,
					BlockNumber: uint64(height),
					TxHash:      txHash,
				}
				stakingEvents = append(stakingEvents, &unbondEvt)
			}
			stakingEvt = nil
		}
	}
	log.Infof("inserting %d staking events", len(stakingEvents))
	return c.db.InsertStakingEvents(stakingEvents)
}

type stakingEventWrapper struct {
	types.CosmosStakingEvent
	srcValidator string
}

func attributesToMap(attributes []abcitypes.EventAttribute) map[string]string {
	m := make(map[string]string, len(attributes))
	for _, a := range attributes {
		m[string(a.Key)] = string(a.Value)
	}
	return m
}

func (c *CosmosIndexer) indexCosmosDelegations(height int64) error {
	log := log.WithField("height", fmt.Sprintf("%d", height))
	var (
		// ctx             = context.Background()
		txSearchResults []*Transaction
		txSearchErr     error
	)

	page := 1
	perPage := 100
	// query := fmt.Sprintf("tx.height=%d AND message.module='staking'", height)
	txSearchResults = make([]*Transaction, 0, 128)
	for {
		// searchResults, err := c.tm.TxSearch(ctx, query, false, &page, &perPage, "asc")
		resp, err := c.lcd.R().SetQueryString(fmt.Sprintf("events=tx.height=19696129&events=message.module='staking'&page=%d&perPage=%d", page, perPage)).Get(c.lcd.BaseURL + "/cosmos/tx/v1beta1/txs?")
		log.Println("Response Info:")
		log.Println("  Error      :", err)
		log.Println("  Status Code:", resp.StatusCode())
		log.Println("  Status     :", resp.Status())
		log.Println("  Proto      :", resp.Proto())
		log.Println("  Time       :", resp.Time())
		log.Println("  Received At:", resp.ReceivedAt())
		log.Println("  Body       :\n", resp)
		log.Println()

		if err != nil {
			txSearchErr = errors.Wrapf(err, "error reading search results height: %d page %d", height, page)
			log.Printf("Error Getting Search Results")
			break
		}

		searchResults := TransactionsResponse{}
		if err = json.Unmarshal(resp.Body(), &searchResults); err != nil {
			log.Errorf("Failed to unmarshal transaction body")
			return errors.Wrapf(err, "error unmarshalling response")
		}

		log.Printf("SEARCH RES: ", searchResults)

		total, err := strconv.Atoi(searchResults.Total)
		if err != nil {
			log.Errorf("Failed to convert total to string")
			return errors.Wrapf(err, "error converting to string")
		}

		txSearchResults = append(txSearchResults, searchResults.Txs...)
		if len(txSearchResults) == total {
			log.Printf("height %d break tx search loop with %d gathered. %d in page %d totalCount %d", height, len(txSearchResults), len(searchResults.Txs), page, total)
			break
		}
		page++
	}

	if txSearchErr != nil {
		log.Printf("Tx Search Err %d", txSearchErr)
		return errors.Wrapf(txSearchErr, "error searching txs block %d", height)
	}

	// for _, t := range txSearchResults {
	// 	if !shouldStoreTx(t.Tx, &t.TxResult) {
	// 		continue
	// 	}
	// 	if err := c.handleStakingTx(height, t.Tx, &t.TxResult); err != nil {
	// 		log.Errorf("error handling staking tx %s: %+v", hashTx(t.Tx), err)
	// 	}
	// }
	return nil
}

func hashTx(bytes []byte) string {
	h := sha256.Sum256(bytes)
	return strings.ToUpper(hex.EncodeToString(h[:]))
}

func parseAmount(in string, decimals uint8) (asset string, amount float64, err error) {
	amt := make([]byte, len(in))
	var i int
	for i = 0; i < len(in); i++ {
		if in[i] < '0' || in[i] > '9' {
			break
		}
		amt[i] = in[i]
	}
	amt = amt[:i]
	asset = in[i:]
	samt := string(amt)
	log.Debugf("%s %s", samt, asset)

	iamt, ok := new(big.Int).SetString(samt, 10)
	if !ok {
		log.Errorf("unable to convert amount %s to big int", samt)
		return
	}

	amount = utils.BigIntToFloat(iamt, uint8(decimals))
	return
}
