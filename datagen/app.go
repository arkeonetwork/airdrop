package datagen

import (
	"context"

	"github.com/ArkeoNetwork/directory/pkg/logging"
	erc20 "github.com/ArkeoNetwork/merkle-drop/contracts"
	"github.com/ArkeoNetwork/merkle-drop/pkg/token_utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AppParams struct {
	EthRPC           string
	FoxGenesisBlock  uint64
	FoxAddressEth    string
	FoxAddressGnosis string
	SnapshotStart    uint64
	SnapshotEnd      uint64
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
	foxAddress := common.HexToAddress(params.FoxAddressEth)

	fox, err := erc20.NewErc20(foxAddress, client)
	if err != nil {
		log.Errorf("failed to create fox %+v", err)
	}
	holders, err := token_utils.GetAllHolders(params.FoxGenesisBlock, blockNumber, 10000, fox)
	if err != nil {
		log.Errorf("failed to get holders of fox %+v", err)
	}

	log.Info(holders)

	return &App{params: params, ethMainnetClient: client}
}
