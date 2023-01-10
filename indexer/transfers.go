package indexer

import (
	"context"
	"errors"

	"github.com/ArkeoNetwork/merkle-drop/contracts/erc20"
	"github.com/ArkeoNetwork/merkle-drop/pkg/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (app *IndexerApp) indexTransfers(startBlock uint64, endBlock uint64, batchSize uint64, token *erc20.Erc20) error {
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
				return errors.New("GetAllTransfers failed with 0 retries")
			}
			continue
		}
		for iter.Next() {

			_, err := app.db.InsertTransfer(types.Transfer{
				From:         iter.Event.From.String(),
				To:           iter.Event.To.String(),
				Value:        iter.Event.Value.String(),
				BlockNumber:  iter.Event.Raw.BlockNumber,
				TxHash:       iter.Event.Raw.TxHash.String(),
				TokenAddress: iter.Event.Raw.Address.String(),
			})
			if err != nil {
				log.Warnf("failed to insert transfer %+v", err)
			}
		}
		currentBlock = end
	}
	return nil
}
