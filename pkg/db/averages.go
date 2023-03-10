package db

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/pkg/errors"
)

type AveragedHolding struct {
	Address string  `db:"account"`
	Holding float64 `db:"avg_hold"`
}

func (d *AirdropDB) FindAveragedBalances(chain, tokenSymbol string) ([]*AveragedHolding, error) {
	conn, err := d.getConnection()
	defer conn.Release()
	if err != nil {
		return nil, errors.Wrapf(err, "error obtaining db connection")
	}

	results := make([]*AveragedHolding, 0, 128)
	if err = pgxscan.Select(context.Background(), conn, &results, sqlFindAveragedBalances, chain, tokenSymbol); err != nil {
		return nil, errors.Wrapf(err, "error querying")
	}
	return results, nil
}
