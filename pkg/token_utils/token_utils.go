package token_utils

import (
	"context"
	"errors"
	"math/big"

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

func GenerateBalanceHistory(allHolders *[]common.Address, transferEvents *[]*erc20.Erc20Transfer, startBlock uint64, endBlock uint64) *map[common.Address]*[]uint64 {
	transfersByAddress := mapTransfersByAddress(transferEvents)

	balanceHistory := map[common.Address]*[]uint64{}
	blockCount := endBlock - startBlock

	for _, address := range *allHolders {
		var balances []uint64 = make([]uint64, blockCount)
		balanceHistory[address] = &balances
		currentBlock := startBlock
		var currentBalance big.Int
		// get all events for this address
		transferEvents := (*transfersByAddress)[address]
		for eventCounter, transferEvent := range *transferEvents {

			for currentBlock < transferEvent.Raw.BlockNumber {
				// fill in users current balance from the last block we saw to
				// this event
				balances[currentBlock-startBlock] = currentBalance.Uint64()
				currentBlock++
			}

			// update the users current balance
			if transferEvent.From == address {
				currentBalance.Sub(&currentBalance, transferEvent.Value)
			} else if transferEvent.To == address {
				currentBalance.Add(&currentBalance, transferEvent.Value)
			} else {
				panic("bad address")
			}

			currentBlock = transferEvent.Raw.BlockNumber

			if eventCounter == len(*transferEvents)-1 {
				// this is the last event for this user, fill forward their balance
				for currentBlock < endBlock {
					balances[currentBlock-startBlock] = currentBalance.Uint64()
					currentBlock++
				}
			}
		}
	}
	return &balanceHistory
}
