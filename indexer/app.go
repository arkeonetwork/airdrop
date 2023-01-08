package indexer

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ArkeoNetwork/directory/pkg/logging"
	"github.com/ArkeoNetwork/merkle-drop/contracts/erc20"
	"github.com/ArkeoNetwork/merkle-drop/pkg/utils/token"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AppParams struct {
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

	var transferEvents *[]*erc20.Erc20Transfer
	// attemp to open jsonFile
	transferJSONFile, err := os.Open("fox_mainnet_transer_events.json")
	if err != nil {
		log.Info("Unable to find transfer events, will re-download")
		transferEvents, err := token.GetAllTransfers(params.FoxGenesisBlock, blockNumber, 1000, fox)
		if err != nil {
			log.Panicf("failed to get transfers of fox %+v", err)
		}

		eventsJSON, err := json.MarshalIndent(transferEvents, "", "  ")
		if err != nil {
			log.Errorf("failed to json %+v", err)
		}

		err = ioutil.WriteFile("fox_mainnet_transer_events.json", eventsJSON, 0644)
		if err != nil {
			log.Errorf("failed to write file %+v", err)
		}
	} else {
		defer transferJSONFile.Close()
		transferJSON, err := ioutil.ReadAll(transferJSONFile)
		if err != nil {
			log.Panic("failed to read JSON")
		} else {
			var transferEventsFromFile []*erc20.Erc20Transfer
			err := json.Unmarshal(transferJSON, &transferEventsFromFile)
			if err != nil {
				log.Panic("failed to unmarshal JSON")
			} else {
				transferEvents = &transferEventsFromFile
			}
		}
	}

	holders := token.GetAllHolders(transferEvents)
	foxMainnetBalanceHistory := token.GenerateBalanceHistory(holders, transferEvents, params.SnapshotStartBlockEth, params.SnapshotEndBlockEth)

	// // deal with FOX LPers and FOX staking on mainnet
	// foxLPAddress := common.HexToAddress(params.FoxLPAddressEth)
	// foxLP, err := erc20.NewErc20(foxLPAddress, client)
	// if err != nil {
	// 	log.Errorf("failed to create foxLP %+v", err)
	// }

	// transferEvents, err = token_utils.GetAllTransfers(params.FoxLPGenesisBlock, blockNumber, 2000, foxLP)
	// if err != nil {
	// 	log.Panicf("failed to get transfers of fox lp %+v", err)
	// }
	// holders = token_utils.GetAllHolders(transferEvents)
	// foxLPBalanceHistory := token_utils.GenerateBalanceHistory(holders, transferEvents, params.SnapshotStartBlockEth, params.SnapshotEndBlockEth)

	// we now need to looking at the FOX LP staking contract.

	weightedBalanceByAddress := token.GetBlockWeigthedAverageBalance(foxMainnetBalanceHistory)
	weightedBalanceByAddressJSON, _ := json.MarshalIndent(weightedBalanceByAddress, "", "  ")
	ioutil.WriteFile("weighted_balances.json", weightedBalanceByAddressJSON, 0644)
	return &App{params: params, ethMainnetClient: client}
}
