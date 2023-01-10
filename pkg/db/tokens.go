package db

import (
	"context"

	"github.com/ArkeoNetwork/merkle-drop/pkg/types"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/pkg/errors"
)

func (d *AirdropDB) FindTokensByChain(chain string) ([]*types.Token, error) {
	conn, err := d.getConnection()
	defer conn.Release()
	if err != nil {
		return nil, errors.Wrapf(err, "error obtaining db connection")
	}
	results := make([]*types.Token, 0, 128)
	if err = pgxscan.Select(context.Background(), conn, &results, sqlFindTokensByChain, chain); err != nil {
		return nil, errors.Wrapf(err, "error scanning")
	}

	return results, nil
}

func (d *AirdropDB) FindAllChainsForTokens() ([]string, error) {
	conn, err := d.getConnection()
	defer conn.Release()
	if err != nil {
		return nil, errors.Wrapf(err, "error obtaining db connection")
	}
	results := make([]string, 10)
	if err = pgxscan.Select(context.Background(), conn, &results, sqlFindAllChains); err != nil {
		return nil, errors.Wrapf(err, "error scanning")
	}

	return results, nil
}

// update token height
func (d *AirdropDB) UpdateTokenHeight(tokenAddress string, height uint64) error {
	conn, err := d.getConnection()
	defer conn.Release()
	if err != nil {
		return errors.Wrapf(err, "error obtaining db connection")
	}
	_, err = conn.Exec(context.Background(), sqlUpdateTokenHeight, height, tokenAddress)
	if err != nil {
		return errors.Wrapf(err, "error updating token height")
	}

	return nil
}
