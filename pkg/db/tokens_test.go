package db

import (
	"testing"
)

func TestFindTokensByChain(t *testing.T) {

	db, err := New(config)
	if err != nil {
		t.Errorf("error getting db: %+v", err)
	}

	tokens, err := db.FindTokensByChain("ETH")

	if err != nil {
		t.Errorf("error finding tokens: %+v", err)
	}

	if len(tokens) == 0 {
		t.Errorf("no tokens found")
	}

}
