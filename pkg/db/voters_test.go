package db

import (
	"testing"

	"github.com/ArkeoNetwork/airdrop/pkg/types"
)

func TestInsertVoters(t *testing.T) {
	db, err := New(config)
	if err != nil {
		t.Errorf("error getting db: %+v", err)
	}

	voters := make([]*types.SnapshotVoter, 1)
	// insert voters
	voter := types.SnapshotVoter{
		Address:       "meaningless-test-address",
	}
	voters = append(voters, &voter)
	err = db.InsertVoters(voters)
	if err != nil {
		t.Errorf("error inserting voters: %+v", err)
	}
}

func TestFindVoterByAddress(t *testing.T) {
	db, err := New(config)
	if err != nil {
		t.Errorf("error getting db: %+v", err)
	}
	hasVoted, err := db.HasAddressParticipatedInProposal("meaningless-test-address")
	if err != nil {
		t.Errorf("error getting voter: %+v", err)
	}
	if hasVoted == false {
		t.Errorf("hasVoted should not be false")
	}
}
