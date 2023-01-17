package db

import (
	"testing"

	"github.com/ArkeoNetwork/airdrop/pkg/types"
)

func TestUpsertTransfer(t *testing.T) {
	db, err := New(config)
	if err != nil {
		t.Errorf("error getting db: %+v", err)
	}

	// upsert transfer
	transfer := types.Transfer{
		TxHash:       "0x123",
		LogIndex:     1,
		TokenAddress: "0xc770eefad204b5180df6a14ee197d99d808ee52d",
		From:         "0x123",
		To:           "0x123",
		Value:        50.55,
		BlockNumber:  1,
	}
	_, err = db.UpsertTransfer(transfer)
	if err != nil {
		t.Errorf("error upserting transfer: %+v", err)
	}
	transfer.Value = 52
	// upsert transfer again
	_, err = db.UpsertTransfer(transfer)
	if err != nil {
		t.Errorf("error upserting transfer: %+v", err)
	}
}

func TestBatchTransfer(t *testing.T) {
	db, err := New(config)
	if err != nil {
		t.Errorf("error getting db: %+v", err)
	}

	transfer1 := types.Transfer{
		TxHash:       "0x123",
		LogIndex:     1,
		TokenAddress: "0xc770eefad204b5180df6a14ee197d99d808ee52d",
		From:         "0x123",
		To:           "0x123",
		Value:        50.55,
		BlockNumber:  1,
	}

	transfer2 := types.Transfer{
		TxHash:       "0x123",
		LogIndex:     2,
		TokenAddress: "0xc770eefad204b5180df6a14ee197d99d808ee52d",
		From:         "0x123",
		To:           "0x123",
		Value:        51.557777666688,
		BlockNumber:  1,
	}

	transfers := []*types.Transfer{&transfer1, &transfer2}
	err = db.UpsertTransferBatch(transfers)
	if err != nil {
		t.Errorf("error upserting transfer: %+v", err)
	}
}

func TestGetBalanceAtBlock(t *testing.T) {
	db, err := New(config)
	if err != nil {
		t.Errorf("error getting db: %+v", err)
	}
	bal, err := db.GetBalanceAtBlock("0xf152a54068c8eddf5d537770985ca8c06ad78abb", 16380085, "0xc770eefad204b5180df6a14ee197d99d808ee52d")
	if err != nil {
		t.Errorf("error getting balance: %+v", err)
	}
	if bal == 0 {
		t.Errorf("balance should not be 0")
	}
}
