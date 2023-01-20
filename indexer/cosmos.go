package indexer

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ArkeoNetwork/airdrop/pkg/db"
	"github.com/ArkeoNetwork/airdrop/pkg/types"
	"github.com/ArkeoNetwork/airdrop/pkg/utils"
	arkutils "github.com/ArkeoNetwork/common/utils"
	"github.com/pkg/errors"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"
)

type CosmosIndexerParams struct {
	DB             arkutils.DBConfig
	Chain          string
	TendermintHost string
	StartHeight    int64
	EndHeight      int64
}

type CosmosIndexer struct {
	db          *db.AirdropDB
	tm          *http.HTTP
	chain       *types.Chain
	startHeight int64
	endHeight   int64
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

	tm, err := arkutils.NewTendermintClient(chain.RpcUrl)
	if err != nil {
		return nil, errors.Wrapf(err, "error creating tendermint client with rpc %s", params.TendermintHost)
	}

	return &CosmosIndexer{db: d, tm: tm, chain: chain, startHeight: params.StartHeight, endHeight: params.EndHeight}, nil
}

// 12940505 first cosmos block on 11/22 UTC 710A86BF2922BA304026102128359EBA941601F2C4B010E2B81D583FAB4A77B1
func (c *CosmosIndexer) IndexDelegators() error {
	startHeight := int64(c.chain.SnapshotStartBlock)
	endHeight := int64(c.chain.SnapshotEndBlock)

	for i := startHeight; i <= endHeight; i++ {
		if err := c.indexDelegations(i); err != nil {
			log.Errorf("error indexing delegations at height %d: %+v", i, err)
		}
	}
	return nil
}

func isStakingTx(tx tmtypes.Tx, txResults *abcitypes.ResponseDeliverTx) bool {
	txHash := hashTx(tx)
	_ = txHash
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
		m := attributesToMap(evt.GetAttributes())
		if i%2 == 0 {
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

			// should be delegate/un/re?
			log.Debug(evt)

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
	log.Infof("have %d staking events", len(stakingEvents))
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

func (c *CosmosIndexer) indexDelegations(height int64) error {
	log := log.WithField("height", fmt.Sprintf("%d", height))
	ctx := context.Background()
	block, err := c.tm.Block(ctx, &height)
	if err != nil {
		return errors.Wrapf(err, "error reading block at %d", height)
	}

	blockResults, err := c.tm.BlockResults(ctx, &height)
	if err != nil {
		return errors.Wrapf(err, "error reading block results at %d", height)
	}

	for i := range block.Block.Txs {
		tx := block.Block.Txs[i]
		txHash := hashTx(tx)
		_ = txHash
		if isStakingTx(tx, blockResults.TxsResults[i]) {
			if err = c.handleStakingTx(height, tx, blockResults.TxsResults[i]); err != nil {
				log.Errorf("error handling staking tx %s: %+v", hashTx(tx), err)
			}
		}
	}
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
