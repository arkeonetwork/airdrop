package indexer

import (
	// "context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ArkeoNetwork/airdrop/pkg/db"
	"github.com/ArkeoNetwork/airdrop/pkg/types"
	"github.com/ArkeoNetwork/airdrop/pkg/utils"
	arkutils "github.com/ArkeoNetwork/common/utils"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	// abcitypes "github.com/tendermint/tendermint/abci/types"
	"io"
	"math/big"
	// "net"
	"net/http"
	"net/url"
	// "net/rpc"
	tmhttp "github.com/tendermint/tendermint/rpc/client/http"
	// coretypes "github.com/tendermint/tendermint/rpc/core/types"
	// tmtypes "github.com/tendermint/tendermint/types"
	"os"
	"strconv"
	"strings"
	"time"
)

type CosmosIndexerParams struct {
	DB    arkutils.DBConfig
	Chain string
}

type CosmosIndexer struct {
	db          *db.AirdropDB
	tm          *tmhttp.HTTP
	lcd         *resty.Client
	chain       *types.Chain
	rpc         string
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

type TxSearchResponse struct {
	Result Result `json:"result"`
}

type Result struct {
	Txs        []*Tx  `json:"txs"`
	TotalCount string `json:"total_count"`
}

type Tx struct {
	Hash     string   `json:"hash"`
	Height   string   `json:"height"`
	TxResult TxResult `json:"tx_result"`
	Tx       string   `json:"tx"`
}

type EventAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Index bool   `json:"index"`
}

type Event struct {
	Type       string           `json:"type"`
	Attributes []EventAttribute `json:"attributes"`
}

type TxResult struct {
	Code      int     `json:"code"`
	Log       string  `json:"log"`
	Info      string  `json:"info"`
	GasWanted string  `json:"gas_wanted"`
	GasUsed   string  `json:"gas_used"`
	Events    []Event `json:"events"`
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
	rpc := chain.RpcUrl
	lcd := resty.New().SetTimeout(10 * time.Second).SetBaseURL(chain.LcdUrl)

	return &CosmosIndexer{db: d, rpc: rpc, lcd: lcd, chain: chain, startHeight: int64(chain.SnapshotStartBlock), endHeight: int64(chain.SnapshotEndBlock)}, nil
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
			break
		}
	}
	return nil
}

func shouldStoreTx(txResults *TxResult) bool {
	for _, evt := range txResults.Events {
		switch evt.Type {
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

func (c *CosmosIndexer) handleStakingTx(height int64, tx string, txResult *TxResult) error {
	log := log.WithField("height", fmt.Sprintf("%d", height))
	txHash := tx
	evtsSequenced := make([]Event, len(txResult.Events))
	evtsSeq := int64(0)
	evtsIndexMap := make(map[int64]int64, 1024)
	for i, evt := range txResult.Events {
		switch evt.Type {
		case "delegate":
			evtsSequenced[evtsSeq] = evt
			evtsIndexMap[evtsSeq] = int64(i)
			evtsSeq++
		case "redelegate":
			evtsSequenced[evtsSeq] = evt
			evtsIndexMap[evtsSeq] = int64(i)
			evtsSeq++
		case "unbond":
			evtsSequenced[evtsSeq] = evt
			evtsIndexMap[evtsSeq] = int64(i)
			evtsSeq++
		case "message":
			m := make(map[string]string, len(evt.Attributes))
			for _, attr := range evt.Attributes {
				m[string(attr.Key)] = string(attr.Value)
			}
			if module, ok := m["module"]; ok && (module == "staking") {
				if delegator, ok := m["sender"]; ok {
					log.Debugf("adding delegate event delegator %s", delegator)
					evtsSequenced[evtsSeq] = evt
					evtsSeq++
				}
			}
		}
	}
	stakingEvents := make([]*types.CosmosStakingEvent, 0, len(evtsSequenced)/2)
	evtsSequenced = evtsSequenced[:evtsSeq]
	var stakingEvt *stakingEventWrapper

	// swap events hack to make sure the first event is a staking event
	if evtsSequenced[0].Type == "message" {
		evtHolder0 := evtsSequenced[0]
		evtHolder1 := evtsSequenced[1]
		evtsSequenced[0] = evtHolder1
		evtsSequenced[1] = evtHolder0
	}
	for i, evt := range evtsSequenced {
		log.Debugf("EVT: %s, Index: %d", evt.Type, i)
		m := attributesToMap(evt.Attributes)
		if evt.Type != "message" {
			// staking event itself: delegate,unbond,redelegate
			_, amount, err := parseAmount(m["amount"], c.chain.Decimals)
			if err != nil {
				log.Errorf("error parsing amount %s: %+v", m["amount"], err)
				return errors.Wrapf(err, "error parsing amount %s", m["amount"])
			}
			var srcValidator, destValidator, delegator string
			var delegatorExists bool
			switch evt.Type {
			case "delegate":
				destValidator = m["validator"]
				delegator, delegatorExists = m["delegator"]
			case "redelegate":
				destValidator = m["destination_validator"]
				srcValidator = m["source_validator"]
			case "unbond":
				srcValidator = m["validator"]
				amount = -amount
			}

			validator := destValidator
			if evt.Type == "unbond" {
				validator = srcValidator
			}

			evtIndex, ok := evtsIndexMap[int64(i)]
			if !ok {
				log.Warnf("no event index found for %d", i)
			} else {
				evtIndex = 0
			}

			stakingEvt = &stakingEventWrapper{
				CosmosStakingEvent: types.CosmosStakingEvent{
					EventType:   evt.Type,
					EventIndex:  evtIndex,
					Validator:   validator,
					Delegator:   delegator,
					Chain:       c.chain.Name,
					Value:       amount,
					BlockNumber: uint64(height),
					TxHash:      txHash,
				},
				srcValidator: srcValidator,
			}

			if delegatorExists && evt.Type == "delegate" {
				stakingEvents = append(stakingEvents, &stakingEvt.CosmosStakingEvent)
			}
		} else {
			// should be message+staking module+spender (delegator address)?
			if stakingEvt == nil {
				return fmt.Errorf("stakingEvt null, event sequencing (programming) issue")
			}
			evtIndex, ok := evtsIndexMap[int64(i)]
			if !ok {
				log.Debugf("no event index found for %d", i)
			} else {
				stakingEvt.EventIndex = evtIndex
			}
			if stakingEvt.Delegator == "" {
				stakingEvt.Delegator = m["sender"]
				stakingEvents = append(stakingEvents, &stakingEvt.CosmosStakingEvent)
			}
			if stakingEvt.EventType == "redelegate" {
				stakingEvents[0].EventType = "delegate"

				unbondEvt := types.CosmosStakingEvent{
					EventType:   "unbond",
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
	log.Debugf("stakingEvt %v", stakingEvt)
	log.Infof("inserting %d staking events", len(stakingEvents))
	if len(stakingEvents) == 0 {
		return errors.New("no staking events to insert")
	}
	totalAmount += len(stakingEvents)
	log.Infof("totalAmount %d", totalAmount)
	return c.db.InsertStakingEvents(stakingEvents)
}

type stakingEventWrapper struct {
	types.CosmosStakingEvent
	srcValidator string
}

func attributesToMap(attributes []EventAttribute) map[string]string {
	m := make(map[string]string, len(attributes))
	for _, a := range attributes {
		m[string(a.Key)] = string(a.Value)
	}
	return m
}
var totalAmount int
func (c *CosmosIndexer) indexCosmosDelegations(height int64) error {
	log := log.WithField("height", fmt.Sprintf("%d", height))
	var (
		txSearchResults []*Tx
		txSearchErr     error
	)

	page := 1
	perPage := 100
	txSearchResults = make([]*Tx, 0, 128)

	for {
		query := url.QueryEscape(fmt.Sprintf("\"tx.height=%d\"", height))
		query = fmt.Sprintf("%s&page=%d&limit=%d", query, page, perPage)
		baseURL := c.rpc
		url := fmt.Sprintf("%s/tx_search?query=%s", baseURL, query)

		log.Infof("Requesting %s", url)

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to make the request: %v", err)
			txSearchErr = errors.Wrapf(err, "Failed to make the request: %v", err)
			break
		}
		defer resp.Body.Close()

		// Check if the request was successful
		if resp.StatusCode != http.StatusOK {
			log.Errorf("Failed to get a successful response: %s", resp.Status)
			txSearchErr = errors.New("Failed to get a successful response")
			break
		}

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Errorf("Failed to read the response body: %v", err)
			txSearchErr = errors.Wrapf(err, "Failed to read the response body: %v", err)
			break
		}

		searchResults := TxSearchResponse{}
		if err = json.Unmarshal(body, &searchResults); err != nil {
			log.Errorf("Failed to unmarshal transaction body")
			txSearchErr = errors.Wrapf(err, "error unmarshalling response")
			break
		}


		totalCount, err := strconv.Atoi(searchResults.Result.TotalCount)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			txSearchErr = errors.Wrapf(err, "error converting string to int")
			break
		}

		if totalCount == 0 {
			log.Debugf("height %d no txs found", height)
			break
		}

		txSearchResults = append(txSearchResults, searchResults.Result.Txs...)
		if len(txSearchResults) == totalCount {
			log.Debugf("height %d break tx search loop with %d gathered. %d in page %d totalCount %s", height, len(txSearchResults), len(searchResults.Result.Txs), page, searchResults.Result.TotalCount)
			break
		}
		page++
	}

	if txSearchErr != nil {
		log.Printf("Tx Search Err %d", txSearchErr)
		return errors.Wrapf(txSearchErr, "error searching txs block %d", height)
	}

	for _, t := range txSearchResults {
		if !shouldStoreTx(&t.TxResult) {
			continue
		}
		if err := c.handleStakingTx(height, t.Hash, &t.TxResult); err != nil {
			log.Errorf("error handling staking tx %s: %+v", hashTx(t.Tx), err)
			return errors.Wrapf(err, "error handling staking tx %s", hashTx(t.Tx))
		}
	}
	return nil
}

func hashTx(tx string) string {
	data := []byte(tx)
	h := sha256.Sum256(data)
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
