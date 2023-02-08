package indexer

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

func (c *CosmosIndexer) IndexOsmoLP(poolName string) error {
	startHeight := int64(c.chain.SnapshotStartBlock)
	endHeight := int64(c.chain.SnapshotEndBlock)

	latestBlock, err := c.tm.Block(context.Background(), nil)
	if err != nil {
		return errors.Wrapf(err, "error getting latest block")
	}

	if endHeight > latestBlock.Block.Height {
		endHeight = latestBlock.Block.Height
	}

	// latest, err := c.db.FindLatestIndexedOsmoLPBlock(c.chain.Name, poolName)
	// if err != nil {
	// 	return errors.Wrapf(err, "error finding latest indexed block")
	// }
	// if latest > startHeight {
	// 	log.Infof("found latest indexed block %d, starting at %d", latest, latest-1)
	// 	startHeight = latest - 1
	// }
	// latest := startHeight
	// _, _ = startHeight, endHeight
	for i := startHeight; i <= endHeight; i++ {
		if err := c.indexOsmoLP(i, poolName); err != nil {
			log.Errorf("error indexing delegations at height %d: %+v", i, err)
		}
		log.Infof("indexed %s lp block %d", poolName, i)
	}
	return nil
}

func (c *CosmosIndexer) indexOsmoLP(height int64, poolName string) error {
	log := log.WithField("height", fmt.Sprintf("%d", height))
	var (
		ctx = context.Background()
		// txSearchResults []*coretypes.ResultTx
		// txSearchErr     error
	)
	_ = log
	// end block events have LP events for chains other than THOR
	blockResults, err := c.tm.BlockResults(ctx, &height)
	if err != nil {
		return errors.Wrapf(err, "error reading search results height %d", height)
	}
	log.Infof("block has %d events", len(blockResults.EndBlockEvents))
	// for _, evt := range blockResults.EndBlockEvents {
	// 	log.Infof("evt %s", evt.Type)
	// }
	page := 1
	perPage := 100
	txSearchResults, err := c.tm.TxSearch(ctx, fmt.Sprintf("tx.height=%d", height), true, &page, &perPage, "asc")
	if err != nil {
		return errors.Wrapf(err, "error reading tx search results height %d", height)
	}
	log.Infof("have %d txResults", len(txSearchResults.Txs))
	for _, txResult := range txSearchResults.Txs {
		txHash := hashTx(txResult.Tx)
		log.Debugf("tx %s", txHash)
		for _, evt := range txResult.TxResult.GetEvents() {
			log.Debugf("evt %s", evt.Type)
			// TODO - if all pools as indicated by notion, decode ibc asset identifiers.
			// otherwise we can whitelist the pools we care about. we'll also need a means
			// to obtain starting balances for each pool.
			switch evt.Type {
			/*
				"attr module: gamm {(height: 7456579)}"
				"attr sender: osmo1sg8r7u5ddqkpha7qv52c6lu82ude88lyf2re46"
				"attr pool_id: 497 {(height: 7456579)}"
				"attr tokens_in: 15635422ibc/46B44899322F3CD854D2D46DEEF881958467CDD4B3B10086DA49296BBED94BED,25672400uosmo {(height: 7456579)}"
			*/
			case "pool_joined":
				log.Infof("tx %s joined pool", txHash)
				for _, attr := range evt.GetAttributes() {
					log.Infof("attr %s: %s", string(attr.Key), string(attr.Value))
				}
			case "pool_exited":
				log.Infof("tx %s exited pool: %v", txHash, evt.GetAttributes())
				for _, attr := range evt.GetAttributes() {
					log.Infof("attr %s: %s", string(attr.Key), string(attr.Value))
				}
			default:
			}
		}
	}
	return nil
}
