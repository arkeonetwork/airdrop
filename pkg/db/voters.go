package db

import (
	"context"
	"strings"

	"github.com/ArkeoNetwork/airdrop/pkg/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

// check if an address has participated in a proposal
func (d *AirdropDB) HasAddressParticipatedInProposal(address string) (bool, error) {
	conn, err := d.getConnection()
	defer conn.Release()
	if err != nil {
		return false, errors.Wrapf(err, "error obtaining db connection")
	}
	var voter types.SnapshotVoter
	if err = selectOne(conn, sqlFindVoterByAddress, &voter, strings.ToLower(address)); err != nil {
		return false, errors.Wrapf(err, "error finding voter")
	}
	// since voter could be an empty struct (because of the `selectOne` implementation),
	// check if `Address` field has a value.
	return voter.Address != "", nil
}

// insert voters
func (d *AirdropDB) InsertVoters(voters []*types.SnapshotVoter) error {
	conn, err := d.getConnection()
	defer conn.Release()
	if err != nil {
		return errors.Wrapf(err, "error obtaining db connection")
	}
	batch := &pgx.Batch{}
	for _, voter := range voters {
		batch.Queue(
			sqlInsertVoters,
			strings.ToLower(voter.Address),
		)
	}
	results := conn.SendBatch(context.Background(), batch)
	err = results.Close()
	if err != nil {
		return errors.Wrap(err, "error inserting voters")
	}
	return nil
}
