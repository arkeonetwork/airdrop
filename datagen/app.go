package datagen

import (
	"context"

	"github.com/ArkeoNetwork/directory/pkg/logging"
	erc20 "github.com/ArkeoNetwork/merkle-drop/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	foxAddress := common.HexToAddress(params.FoxAddressEth)

	fox, err := erc20.NewErc20(foxAddress, client)
	if err != nil {
		log.Errorf("failed to create fox %+v", err)
	}
	snapshotEnd := uint64(params.FoxGenesisBlock + 1500000)
	//snapshotEnd := uint64(params.SnapshotEnd)

	filterOpts := bind.FilterOpts{
		Start:   uint64(params.FoxGenesisBlock),
		End:     &snapshotEnd,
		Context: context.Background(),
	}
	iter, err := fox.FilterTransfer(&filterOpts, nil, nil)

	//logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Errorf("failed to get logs from eth RPC client %+v", err)
	}

	for iter.Next() {
		log.Info(iter.Event.From, iter.Event.To)
	}

	return &App{params: params, ethMainnetClient: client}
}
