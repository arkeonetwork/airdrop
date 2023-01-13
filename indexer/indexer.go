package indexer

import (
	"context"
	"fmt"

	"github.com/ArkeoNetwork/airdrop/contracts/erc20"
	"github.com/ArkeoNetwork/airdrop/pkg/db"
	"github.com/ArkeoNetwork/directory/pkg/logging"
<<<<<<< HEAD
=======
	"github.com/ArkeoNetwork/merkle-drop/contracts/erc20"
	"github.com/ArkeoNetwork/merkle-drop/pkg/db"
>>>>>>> f98d274 (adds multichain functionality)
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IndexerAppParams struct {
<<<<<<< HEAD
	SnapshotStart uint64
	SnapshotEnd   uint64
=======
	SnapshotStart         uint64
	SnapshotEnd           uint64
	SnapshotStartBlockEth uint64
	SnapshotEndBlockEth   uint64
>>>>>>> f98d274 (adds multichain functionality)
	db.DBConfig
}

type IndexerApp struct {
	params IndexerAppParams
	db     *db.AirdropDB
}

var log = logging.WithoutFields()

func NewIndexer(params IndexerAppParams) *IndexerApp {
	d, err := db.New(params.DBConfig)
	if err != nil {
		panic(fmt.Sprintf("error connecting to the db: %+v", err))
	}
	return &IndexerApp{params: params, db: d}
}

func (app *IndexerApp) Start() {
	// find all chains we care about
	chains, err := app.db.FindAllChains()
	if err != nil {
		panic(fmt.Sprintf("unbale to find chains for tokens: %+v", err))
	}

	// todo: kick off new go-routine for each chain
	for _, chain := range chains {
		client, err := ethclient.Dial(chain.RpcUrl)
		if err != nil {
			panic(fmt.Sprintf("failed to connect to eth RPC client %+v", err))
		}

		blockNumber, err := client.BlockNumber(context.Background())
		if err != nil {
			panic(fmt.Sprintf("failed to get current block number from RPC client %+v", err))
		}

		log.Infof("Connected to client for %s. Current block %d", chain.Name, blockNumber)
<<<<<<< HEAD
		snapshotEnd := blockNumber
		if blockNumber > chain.SnapshotEndBlock {
			snapshotEnd = chain.SnapshotEndBlock
=======
		snapshotEndEth := blockNumber
		if blockNumber > app.params.SnapshotEndBlockEth {
			snapshotEndEth = app.params.SnapshotEndBlockEth
>>>>>>> f98d274 (adds multichain functionality)
		}

		// get the tokens for each chain
		tokens, err := app.db.FindTokensByChain(chain.Name)
		if err != nil {
			panic(fmt.Sprintf("unbale to find tokens for chain: %+v", err))
		}

		// iterate tokens array and get the transfers for each token
		for _, token := range tokens {
			// determine if the token has been synced to a differnt block
			startBlock := token.GenesisBlock
			if token.Height > startBlock {
				startBlock = token.Height
			}

			// get the transfers for each token
<<<<<<< HEAD
			log.Infof("Getting transfers for token: %s from block: %d to block: %d", token.Name, startBlock, snapshotEnd)
=======
			log.Infof("Getting transfers for token: %s from block: %d to block: %d", token.Name, app.params.SnapshotStartBlockEth, snapshotEndEth)
>>>>>>> f98d274 (adds multichain functionality)
			tokenContract, err := erc20.NewErc20(common.HexToAddress(token.Address), client)
			if err != nil {
				panic(fmt.Sprintf("unbale to get token contract for token: %+v", err))
			}

<<<<<<< HEAD
			err = app.IndexTransfers(startBlock, snapshotEnd, 1000, token.Address, tokenContract)
=======
			err = app.IndexTransfers(startBlock, snapshotEndEth, 1000, token.Address, tokenContract)
>>>>>>> f98d274 (adds multichain functionality)
			if err != nil {
				panic(fmt.Sprintf("unbale to get transfers for token: %+v", err))
			}
		}
	}

}
