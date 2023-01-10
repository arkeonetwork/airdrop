package indexer

import (
	"context"
	"fmt"

	"github.com/ArkeoNetwork/directory/pkg/logging"
	"github.com/ArkeoNetwork/merkle-drop/pkg/db"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IndexerAppParams struct {
	EthRPC                 string
	FoxGenesisBlock        uint64
	FoxLPGenesisBlock      uint64
	FoxAddressEth          string
	FoxLPAddressEth        string
	FoxStakingAddressEth   string
	FoxStakingGenesisBlock uint64
	FoxAddressGnosis       string
	SnapshotStart          uint64
	SnapshotEnd            uint64
	SnapshotStartBlockEth  uint64
	SnapshotEndBlockEth    uint64
	db.DBConfig
}

type IndexerApp struct {
	params           IndexerAppParams
	ethMainnetClient *ethclient.Client
	db               *db.AirdropDB
}

var log = logging.WithoutFields()

func NewIndexer(params IndexerAppParams) *IndexerApp {
	client, err := ethclient.Dial(params.EthRPC)
	if err != nil {
		panic(fmt.Sprintf("failed to connet to eth RPC client %+v", err))
	}
	_, err = client.BlockNumber(context.Background())
	if err != nil {
		panic(fmt.Sprintf("failed to get current block number from eth RPC client %+v", err))
	}

	d, err := db.New(params.DBConfig)
	if err != nil {
		panic(fmt.Sprintf("error connecting to the db: %+v", err))
	}
	return &IndexerApp{params: params, ethMainnetClient: client, db: d}
}

func (app *IndexerApp) Start() {
	blockNumber, err := app.ethMainnetClient.BlockNumber(context.Background())
	if err != nil {
		panic(fmt.Sprintf("failed to get current block number from eth RPC client %+v", err))
	}

	log.Infof("Connected to eth mainnet client. Current block %d", blockNumber)
	snapshotEndEth := blockNumber
	if blockNumber > app.params.SnapshotEndBlockEth {
		snapshotEndEth = app.params.SnapshotEndBlockEth
	}

	// find all chains we care about
	chains, err := app.db.FindAllChainsForTokens()
	if err != nil {
		panic(fmt.Sprintf("unbale to find chains for tokens: %+v", err))
	}

	for _, chain := range chains {
		// skip all chains but eth for tesing
		if chain != "ETH" {
			continue
		}

		// get the tokens for each chain
		tokens, err := app.db.FindTokensByChain(chain)
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
			log.Infof("Getting transfers for token: %s from block: %d to block: %d", token.Name, app.params.SnapshotStartBlockEth, snapshotEndEth)
			err = app.IndexTransfers(startBlock, snapshotEndEth, 1000, token.Address)
			if err != nil {
				panic(fmt.Sprintf("unbale to get transfers for token: %+v", err))
			}
		}
	}

}
