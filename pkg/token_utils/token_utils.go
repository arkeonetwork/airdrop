package token_utils

import (
	"context"

	"github.com/ArkeoNetwork/directory/pkg/logging"

	erc20 "github.com/ArkeoNetwork/merkle-drop/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var log = logging.WithoutFields()

func GetAllHolders(startBlock uint64, endBlock uint64, batchSize uint64, token *erc20.Erc20) (*[]common.Address, error) {
	isHolder := map[common.Address]bool{}
	currentBlock := startBlock
	retryCount := 20
	for currentBlock < endBlock && retryCount > 0 {
		end := currentBlock + batchSize
		filterOpts := bind.FilterOpts{
			Start:   currentBlock,
			End:     &end,
			Context: context.Background(),
		}
		iter, err := token.FilterTransfer(&filterOpts, nil, nil)
		if err != nil {
			// retry?
			log.Errorf("failed to get transfer events for block %+v", err)
			retryCount--
			continue
		}
		for iter.Next() {
			isHolder[iter.Event.To] = true
		}

		currentBlock = end
	}

	holders := make([]common.Address, len(isHolder))

	i := 0
	for k := range isHolder {
		holders[i] = k
		i++
	}

	return &holders, nil
}
