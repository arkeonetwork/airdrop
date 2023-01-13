package db

import (
	"context"

	"github.com/ArkeoNetwork/airdrop/pkg/types"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/pkg/errors"
)

// findallchains - queries chains table for all chains
func (d *AirdropDB) FindAllChains() ([]*types.Chain, error) {
	conn, err := d.getConnection()
	defer conn.Release()
	if err != nil {
		return nil, errors.Wrapf(err, "error obtaining db connection")
	}
	results := make([]*types.Chain, 0, 10)
	if err = pgxscan.Select(context.Background(), conn, &results, sqlFindAllChains); err != nil {
		return nil, errors.Wrapf(err, "error scanning")
	}
	return results, nil
}
