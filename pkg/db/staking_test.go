package db

import (
	"testing"
)

func TestFindStakingContracts(t *testing.T) {

	db, err := New(config)
	if err != nil {
		t.Errorf("error getting db: %+v", err)
	}

	contracts, err := db.FindStakingContracts()

	if err != nil {
		t.Errorf("error finding staking contracts: %+v", err)
	}

	if len(contracts) == 0 {
		t.Errorf("no staking contracts found")
	}

}

func TestFindStakingContractsByName(t *testing.T) {

	db, err := New(config)
	if err != nil {
		t.Errorf("error getting db: %+v", err)
	}

	contracts, err := db.FindStakingContractsByName("stakingrewards")

	if err != nil {
		t.Errorf("error finding staking contracts: %+v", err)
	}

	if len(contracts) == 0 {
		t.Errorf("no staking contracts found")
	}
}
