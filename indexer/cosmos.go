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
					log.Infof("adding delegate event delegator %s", delegator)
					// log.Infof("setting module_staking_delegator to %s", delegator)
					// attribMap["module_staking_delegator"] = delegator
					evtsSequenced[evtsSeq] = evt
					evtsIndexMap[evtsSeq] = int64(i)
					evtsSeq++
				}
			}
		}
	}
	stakingEvents := make([]*types.CosmosStakingEvent, 0, len(evtsSequenced)/2)
	evtsSequenced = evtsSequenced[:evtsSeq]
	var stakingEvt *types.CosmosStakingEvent
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

			stakingEvt = &types.CosmosStakingEvent{
				EventType:   evt.GetType(),
				Validator:   validator,
				Chain:       c.chain.Name,
				Value:       amount,
				BlockNumber: uint64(height),
				TxHash:      txHash,
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
			stakingEvents = append(stakingEvents, stakingEvt)
			stakingEvt = nil
		}
	}
	log.Infof("have %d staking events", len(stakingEvents))
	return c.db.InsertStakingEvents(stakingEvents)
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
	log.Infof("%s %s", samt, asset)

	iamt, ok := new(big.Int).SetString(samt, 10)
	if !ok {
		log.Errorf("unable to convert amount %s to big int", samt)
		return
	}

	amount = utils.BigIntToFloat(iamt, uint8(decimals))
	return
}

/*
type txResponseWrapper interface {
	getResponseDeliverTx() *abcitypes.ResponseDeliverTx
	getTx() tmtypes.Tx
}
type baseTxResponseWrapper struct {
	tx       tmtypes.Tx
	txResult *abcitypes.ResponseDeliverTx
}

func (w baseTxResponseWrapper) getTx() tmtypes.Tx {
	return w.tx
}

func (w baseTxResponseWrapper) getResponseDeliverTx() *abcitypes.ResponseDeliverTx {
	return w.txResult
}

type delegateTxResponse struct {
	baseTxResponseWrapper
}
type undelegateTxResponse struct {
	baseTxResponseWrapper
}
type redelegateTxResponse struct {
	baseTxResponseWrapper
}
type unbondTxResponse struct {
	baseTxResponseWrapper
}

func (c *CosmosIndexer) handleStakingEvents(height int64, txResponses map[string]txResponseWrapper) error {
	stakingEvents := make([]*types.CosmosStakingEvent, 0, len(txResponses))
	for _, txw := range txResponses {
		dumpTxResult(txw.getResponseDeliverTx())
		var err error
		var stakingEvent *types.CosmosStakingEvent
		switch v := txw.(type) {
		case delegateTxResponse:
			stakingEvent, err = c.handleDelegate(height, v)
		case undelegateTxResponse:
			err = handleUndelegate(height, v)
		case redelegateTxResponse:
			err = handleRedelegate(height, v)
		case unbondTxResponse:
			err = handleUnbond(height, v)
		}
		if err != nil {
			log.Errorf("error handling %s staking event: %+v", txw.getResponseDeliverTx().GetInfo(), err)
			continue
		}
		if stakingEvent == nil {
			log.Warnf("no staking event produced for %s", hashTx(txw.getTx()))
			continue
		}
		stakingEvents = append(stakingEvents, stakingEvent)
	}
	if len(stakingEvents) > 0 {
		log.Infof("found %d staking events to write at height %d", len(stakingEvents), height)
		if err := c.db.InsertStakingEvents(stakingEvents); err != nil {
			return errors.Wrapf(err, "error inserting staking events at height %d, abort", height)
		}
	}

	return nil
}

func (c *CosmosIndexer) handleDelegate(height int64, d delegateTxResponse) (*types.CosmosStakingEvent, error) {
	tx := d.tx
	txHash := hashTx(tx)
	log.Infof("tx %s has delegate event", txHash)
	txResult := d.txResult

	attribMap := make(map[string]string, 64)
	evtsSequenced := make([]abcitypes.Event, len(txResult.GetEvents()))
	evtsSeq := 0
	for _, evt := range txResult.GetEvents() {
		switch evt.GetType() {
		case "delegate":
			evtsSequenced[evtsSeq] = evt
			evtsSeq++
			// for _, attrib := range evt.GetAttributes() {
			// 	attribMap[string(attrib.GetKey())] = string(attrib.GetValue())
			// }
		// case "transfer":
		// 	for _, attrib := range evt.GetAttributes() {
		// 		attribMap[string(attrib.GetKey())] = string(attrib.GetValue())
		// 	}
		case "message":
			m := make(map[string]string, len(evt.GetAttributes()))
			for _, attr := range evt.GetAttributes() {
				m[string(attr.GetKey())] = string(attr.GetValue())
			}
			if module, ok := m["module"]; ok && module == "staking" {
				if delegator, ok := m["sender"]; ok {
					log.Infof("adding delegate event delegator %s", delegator)
					// log.Infof("setting module_staking_delegator to %s", delegator)
					// attribMap["module_staking_delegator"] = delegator
					evtsSequenced[evtsSeq] = evt
					evtsSeq++
				}
			}
		// case "tx":
		// 	for _, attrib := range evt.GetAttributes() {
		// 		if string(attrib.GetKey()) == "acc_seq" {
		// 			spl := strings.Split(string(attrib.GetValue()), "/")
		// 			attribMap["delegator"] = spl[0]
		// 			break
		// 		}
		// 		attribMap[string(attrib.GetKey())] = string(attrib.GetValue())
		// 	}
		default:
		}
	}
	if len(attribMap) < 1 {
		return nil, fmt.Errorf("no attributes for delegate event, height: %d, tx: %s", height, txHash)
	}
	_, amt, err := parseAmount(attribMap["amount"], c.chain.Decimals)
	if err != nil {
		log.Errorf("error parsing amount %s: %+v", attribMap["amount"], err)
	}
	cse := &types.CosmosStakingEvent{
		EventType:   "delegate",
		Chain:       c.chain.Name,
		BlockNumber: uint64(height),
		TxHash:      txHash,
		Delegator:   attribMap["delegator"],
		Validator:   attribMap["validator"],
		Value:       amt,
	}

	return cse, nil
}
*/

// func dumpTxResult(r *abcitypes.ResponseDeliverTx) {
// 	uniq := make(map[string]string, 256)
// 	for _, evt := range r.Events {
// 		// log.Infof("evt: %s", evt.GetType())
// 		for _, attrib := range evt.GetAttributes() {
// 			// log.Infof("%s : %s", attrib.GetKey(), attrib.GetValue())
// 			uniq[string(attrib.GetKey())] = string(attrib.GetValue())
// 		}
// 	}
// 	for k, v := range uniq {
// 		log.Infof("%s : %s", k, v)
// 	}
// }
