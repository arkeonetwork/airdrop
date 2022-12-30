package token_utils

import (
	"context"
	"errors"
	"math/big"
	"sync"

	"github.com/ArkeoNetwork/directory/pkg/logging"

	erc20 "github.com/ArkeoNetwork/merkle-drop/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var log = logging.WithoutFields()

func GetAllTransfers(startBlock uint64, endBlock uint64, batchSize uint64, token *erc20.Erc20) (*[]*erc20.Erc20Transfer, error) {
	transfers := []*erc20.Erc20Transfer{}
	currentBlock := startBlock
	retryCount := 20
	for currentBlock < endBlock {
		end := currentBlock + batchSize
		filterOpts := bind.FilterOpts{
			Start:   currentBlock,
			End:     &end,
			Context: context.Background(),
		}
		iter, err := token.FilterTransfer(&filterOpts, nil, nil)
		if err != nil {
			log.Errorf("failed to get transfer events for block %+v retring", err)
			retryCount--
			if retryCount < 0 {
				return nil, errors.New("GetAllTransfers failed with 0 retries")
			}
			continue
		}
		for iter.Next() {
			transfers = append(transfers, iter.Event)
		}
		currentBlock = end
	}
	return &transfers, nil
}

func GetAllHolders(transferEvents *[]*erc20.Erc20Transfer) *[]common.Address {
	isHolder := map[common.Address]bool{}
	for _, transfer := range *transferEvents {
		isHolder[transfer.To] = true
	}
	holders := make([]common.Address, len(isHolder))

	i := 0
	for k := range isHolder {
		holders[i] = k
		i++
	}

	return &holders
}

func mapTransfersByAddress(transferEvents *[]*erc20.Erc20Transfer) *map[common.Address]*[]*erc20.Erc20Transfer {
	transfersByAddress := map[common.Address]*[]*erc20.Erc20Transfer{}
	for _, transfer := range *transferEvents {
		addTransferToAddressMap(&transfersByAddress, transfer, transfer.From)
		addTransferToAddressMap(&transfersByAddress, transfer, transfer.To)
	}
	return &transfersByAddress
}

func addTransferToAddressMap(transfersByAddress *map[common.Address]*[]*erc20.Erc20Transfer, transfer *erc20.Erc20Transfer, address common.Address) {
	transfersByAddressLocal := *transfersByAddress
	transfers, exists := transfersByAddressLocal[address]
	if !exists {
		transfers := []*erc20.Erc20Transfer{transfer}
		transfersByAddressLocal[address] = &transfers
	} else {
		updatedTransfers := append(*transfers, transfer)
		transfersByAddressLocal[address] = &updatedTransfers
	}
}

func GetBalancesAtBlock(allHolders *[]common.Address, transferEvents *[]*erc20.Erc20Transfer, block uint64) *map[common.Address]*big.Int {
	transfersByAddress := mapTransfersByAddress(transferEvents)
	balancesAtBlockByAddress := map[common.Address]*big.Int{}
	for _, address := range *allHolders {
		currentBalance := big.NewInt(0)
		// get all events for this address
		transferEvents := (*transfersByAddress)[address]
		for _, transferEvent := range *transferEvents {
			if transferEvent.Raw.BlockNumber > block {
				break
			}

			// update the users current balance
			if transferEvent.From == address {
				currentBalance.Sub(currentBalance, transferEvent.Value)
			} else if transferEvent.To == address {
				currentBalance.Add(currentBalance, transferEvent.Value)
			} else {
				panic("bad address")
			}

		}
		balancesAtBlockByAddress[address] = currentBalance
	}
	return &balancesAtBlockByAddress
}

func GenerateBalanceHistory(
	allHolders *[]common.Address,
	transferEvents *[]*erc20.Erc20Transfer,
	startingBalances *map[common.Address]*big.Int,
	startBlock uint64,
	endBlock uint64,
) *map[common.Address]*[]*big.Int {
	transfersByAddress := mapTransfersByAddress(transferEvents)

	balanceHistory := map[common.Address]*[]*big.Int{}
	blockCount := endBlock - startBlock

	for i, address := range *allHolders {
		var balances []*big.Int = make([]*big.Int, blockCount)
		balanceHistory[address] = &balances
		currentBlock := startBlock
		var currentBalance *big.Int = new(big.Int).Set((*startingBalances)[address])

		// get all events for this address
		transferEvents := (*transfersByAddress)[address]
		log.Debugf("Processing %d of %d : %s with %d transfer events", i, len(*allHolders), address.String(), len(*transferEvents))
		for eventCounter, transferEvent := range *transferEvents {
			if transferEvent.Raw.BlockNumber > endBlock {
				break
			}

			for currentBlock < transferEvent.Raw.BlockNumber {
				// fill in users current balance from the last block we saw to
				// this event
				balances[currentBlock-startBlock] = currentBalance
				currentBlock++
			}

			// update the users current balance
			if transferEvent.From == address {
				currentBalance = new(big.Int).Set(currentBalance)
				currentBalance = currentBalance.Sub(currentBalance, transferEvent.Value)
			} else if transferEvent.To == address {
				currentBalance = new(big.Int).Set(currentBalance)
				currentBalance = currentBalance.Add(currentBalance, transferEvent.Value)
			} else {
				panic("bad address")
			}

			if eventCounter == len(*transferEvents)-1 {
				// this is the last event for this user, fill forward their balance
				for currentBlock < endBlock {
					balances[currentBlock-startBlock] = currentBalance
					currentBlock++
				}
			}
		}
	}
	return &balanceHistory
}

func GetBlockWeigthedAverageBalance(balanceHistory *map[common.Address]*[]*big.Int) *map[common.Address]*big.Int {
	averageBalanceByAddress := map[common.Address]*big.Int{}
	mutex := sync.Mutex{}
	var waitGroup sync.WaitGroup
	for addressKey, balancesVal := range *balanceHistory {
		waitGroup.Add(1)
		go func(address common.Address, balances []*big.Int, avgBalanceByAddress *map[common.Address]*big.Int) {
			defer waitGroup.Done()
			balance := getAverageBalance(&balances)
			mutex.Lock() // go cannot handle concurrent writes to different keys in the map.
			(*avgBalanceByAddress)[address] = balance
			mutex.Unlock()
		}(addressKey, *balancesVal, &averageBalanceByAddress)
	}
	waitGroup.Wait()
	return &averageBalanceByAddress
}

func getAverageBalance(balances *[]*big.Int) *big.Int {
	balanceSum := big.NewInt(0)
	for _, balance := range *balances {
		balanceSum.Add(balanceSum, balance)
	}
	return balanceSum.Div(balanceSum, big.NewInt(int64(len(*balances))))
}
