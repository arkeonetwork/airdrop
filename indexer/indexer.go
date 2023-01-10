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
	db               *db.DirectoryDB
}

var log = logging.WithoutFields()

func NewIndexer(params IndexerAppParams) *IndexerApp {
	client, err := ethclient.Dial(params.EthRPC)
	if err != nil {
		log.Panicf("failed to connet to eth RPC client %+v", err)
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

func (app *IndexerApp) start() {
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
		tokens, err := app.db.FindAllTokensForChain(chain)
		if err != nil {
			panic(fmt.Sprintf("unbale to find tokens for chain: %+v", err))
		}

		// iterate tokens array and get the transfers for each token
		for _, token := range tokens {
			// get the transfers for each token
			transferEvents, err := token_utils.GetAllTransfers(app.params.SnapshotStartBlockEth, snapshotEndEth, 10000, token)
			if err != nil {
				panic(fmt.Sprintf("unbale to get transfers for token: %+v", err))
			}
			// get the holders for each token
			holders := token_utils.GetAllHolders(transferEvents)
			// iterate holders array and get the balances for each holder
			for _, holder := range *holders {
				balance, err := token.BalanceOf(nil, holder)
				if err != nil {
					panic(fmt.Sprintf("unbale to get balance for holder: %+v", err))
				}
				// save the balances for each holder
				err = app.db.SaveBalance(chain, token, holder, balance)
				if err != nil {
					panic(fmt.Sprintf("unbale to save balance for holder: %+v", err))
				}
			}
		}
	}

}
