package db

import (
	"context"

	"github.com/ArkeoNetwork/merkle-drop/pkg/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (d *AirdropDB) UpsertTransfer(tx types.Transfer) (*Entity, error) {
	conn, err := d.getConnection()
	defer conn.Release()
	if err != nil {
		return nil, errors.Wrapf(err, "error obtaining db connection")
	}
	return upsert(conn, sqlUpsertTransferEvent, tx.TxHash, tx.LogIndex, tx.TokenAddress, tx.From, tx.To, tx.Value, tx.BlockNumber)
}

func (d *AirdropDB) UpsertTransferBatch(transfers []*types.Transfer) error {
	conn, err := d.getConnection()
	defer conn.Release()
	if err != nil {
		return errors.Wrapf(err, "error obtaining db connection")
	}
	x := make([]interface{}, len(transfers))
	for i, _ := range transfers {
		x[i] = transfers[i]
	}

	batch := &pgx.Batch{}
	for _, transfer := range transfers {
		batch.Queue(
			sqlUpsertTransferEvent,
			transfer.TxHash,
			transfer.LogIndex,
			transfer.TokenAddress,
			transfer.From,
			transfer.To,
			transfer.Value,
			transfer.BlockNumber)
	}
	results := conn.SendBatch(context.Background(), batch)
	err = results.Close()
	if err != nil {
		return errors.Wrap(err, "error executing batch")
	}
	return nil

}
