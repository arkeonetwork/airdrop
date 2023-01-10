package indexer

import (
	"context"
	"errors"
	"math"
	"math/big"

	"github.com/ArkeoNetwork/merkle-drop/contracts/erc20"
	"github.com/ArkeoNetwork/merkle-drop/pkg/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (app *IndexerApp) IndexTransfers(startBlock uint64, endBlock uint64, batchSize uint64, tokenAddress string) error {
	token, err := erc20.NewErc20(common.HexToAddress(tokenAddress), app.ethMainnetClient)
	if err != nil {
		log.Errorf("failed to create token %+v", err)
		return err
	}
	decimals, err := token.Decimals(nil)
	if err != nil {
		log.Errorf("failed to get token decimals %+v", err)
		return err
	}
	name, err := token.Name(nil)
	if err != nil {
		log.Errorf("failed to get token name %+v", err)
		return err
	}
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
		// add each transfer to array
		transfers := []*types.Transfer{}
		for iter.Next() {
			// todo: batch insert
			transferValue := new(big.Float).SetInt(iter.Event.Value)
			transferValue.Quo(transferValue, big.NewFloat(float64(math.Pow10(int(decimals)))))
			transferValueDecimal, _ := transferValue.Float64()
			transfers = append(transfers,
				&types.Transfer{
					From:         iter.Event.From.String(),
					To:           iter.Event.To.String(),
					Value:        transferValueDecimal,
					BlockNumber:  iter.Event.Raw.BlockNumber,
					TxHash:       iter.Event.Raw.TxHash.String(),
					TokenAddress: tokenAddress,
				})
		}
		err = app.db.UpdateTokenHeight(tokenAddress, end)
		if err != nil {
			log.Warnf("failed to update token height %+v", err)
		}
		currentBlock = end

		if len(transfers) == 0 {
			continue
		}

		err = app.db.UpsertTransferBatch(transfers)
		if err != nil {
			log.Errorf("failed to upsert transfer batch %+v", err)
			return err
		}
		log.Debugf("%s: updated transfers for blocks through %d with %d transfers", name, end, len(transfers))
	}
	return nil
}
