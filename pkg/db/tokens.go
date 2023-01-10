package db

import (
	"context"

	"github.com/ArkeoNetwork/merkle-drop/pkg/types"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/pkg/errors"
)

func (d *DirectoryDB) FindTokensByChain(chain string) ([]*types.Token, error) {
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

func (d *DirectoryDB) FindAllChainsForTokens() ([]string, error) {
	conn, err := d.getConnection()
	if err != nil {
		return nil, errors.Wrapf(err, "error obtaining db connection")
	}
	results := make([]string, 10)
	if err = pgxscan.Select(context.Background(), conn, &results, sqlFindAllChains); err != nil {
		return nil, errors.Wrapf(err, "error scanning")
	}

	return results, nil
}
