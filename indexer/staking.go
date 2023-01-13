package indexer

import (
	"context"

	"github.com/ArkeoNetwork/airdrop/contracts/stakingrewards"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

// indexLPStaking indexes the LP staking events
func (app *IndexerApp) IndexLPStaking() error {
	// get all staking contracts
	stakingContracts, err := app.db.FindStakingContracts()
	if err != nil {
		return errors.Wrap(err, "error finding all staking contracts")
	}
	// for each staking contract
	for _, stakingContract := range stakingContracts {
		log.Info("indexing staking contract: ", stakingContract.ContractName)
		// get chain info and create rpc client
		chain, err := app.db.FindChain(stakingContract.Chain)
		if err != nil {
			return errors.Wrap(err, "error finding chain for staking contract")
		}
		client, err := ethclient.Dial(chain.RpcUrl)
		if err != nil {
			return errors.Wrapf(err, "failed to connect to eth RPC client %s", chain.RpcUrl)
		}
		// create staking contract
		staking, err := stakingrewards.NewStakingrewards(common.HexToAddress(stakingContract.Address), client)
		if err != nil {
			return errors.Wrap(err, "error creating staking contract")
		}
		// get current block number
		startBlock := stakingContract.Height
		if startBlock < stakingContract.GenesisBlock {
			startBlock = stakingContract.GenesisBlock
		}

		blockNumber, err := client.BlockNumber(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed to get current block number from RPC client")
		}
		endBlock := chain.SnapshotEndBlock
		if blockNumber < endBlock {
			endBlock = blockNumber
		}
		log.Infof("Connected to client for %s. Current block %d Indexing staking events from block %d", chain.Name, blockNumber, startBlock)
	}
}
