package db

import (
	"context"

	"github.com/ArkeoNetwork/airdrop/pkg/types"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/pkg/errors"
)

// findStakingContracts
func (d *AirdropDB) FindStakingContracts() ([]*types.Staking, error) {
	conn, err := d.getConnection()
	defer conn.Release()
	if err != nil {
		return nil, errors.Wrapf(err, "error obtaining db connection")
	}
	results := make([]*types.Staking, 0, 10)
	if err = pgxscan.Get(context.Background(), conn, results, sqlFindAllStakingContracts); err != nil {
		return nil, errors.Wrapf(err, "error scanning")
	}
	return results, nil
}
