package datagen

import (
	"context"
	"math/big"

	"github.com/ArkeoNetwork/directory/pkg/logging"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AppParams struct {
	EthRPC           string
	FoxGenesisBlock  int
	FoxAddressEth    string
	FoxAddressGnosis string
	SnapshotStart    int
	SnapshotEnd      int
}

type App struct {
	params           AppParams
	ethMainnetClient *ethclient.Client
}

var log = logging.WithoutFields()

func NewApp(params AppParams) *App {
	client, err := ethclient.Dial(params.EthRPC)
	if err != nil {
		log.Panicf("failed to connet to eth RPC client %+v", err)
	}
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Panicf("failed to get current block number from eth RPC client %+v", err)
	}

	log.Infof("Connected to eth mainnet client. Current block %d", blockNumber)

	// query := ethereum.FilterQuery{
	// 	FromBlock: big.NewInt(2394201),
	// 	ToBlock:   big.NewInt(2394201),
	// 	Addresses: []common.Address{
	// 	  contractAddress,
	// 	},
	//   }

	foxAddress := common.HexToAddress(params.FoxAddressEth)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(params.FoxGenesisBlock)),
		ToBlock:   big.NewInt(int64(params.SnapshotEnd)),
		Addresses: []common.Address{
			foxAddress,
		},
	}

	client.FilterLogs(context.Background(), query)
	// we should now be able to get all the holders of FOX that could potentially have a balance during the airdrop

	return &App{params: params, ethMainnetClient: client}
}
