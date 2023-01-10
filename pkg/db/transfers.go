package db

import (
	"github.com/ArkeoNetwork/merkle-drop/pkg/types"
	"github.com/pkg/errors"
)

func (d *AirdropDB) InsertTransfer(tx types.Transfer) (*Entity, error) {
	conn, err := d.getConnection()
	defer conn.Release()
	if err != nil {
		return nil, errors.Wrapf(err, "error obtaining db connection")
	}
	return insert(conn, sqlInsertTransferEvent, tx.TxHash, tx.TokenAddress, tx.From, tx.To, tx.Value, tx.BlockNumber)
}
