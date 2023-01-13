package token

import (
	"math/big"
	"testing"

	"github.com/ArkeoNetwork/airdrop/contracts/erc20"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestGenerateBalanceHistory(t *testing.T) {
	addressA := common.HexToAddress("0xc770eefad204b5180df6a14ee197d99d808ee52d")
	addressB := common.HexToAddress("0x21a42669643f45bc0e086b8fc2ed70c23d67509d")
	addressZero := common.HexToAddress("0x0")

	allHolders := []common.Address{addressA, addressB}

	startBlock := 100
	endBlock := 200

	transferOne := erc20.Erc20Transfer{
		From:  addressZero,
		To:    addressA,
		Value: big.NewInt(500),
		Raw: types.Log{
			BlockNumber: 125,
		},
	}

	transferTwo := erc20.Erc20Transfer{
		From:  addressA,
		To:    addressB,
		Value: big.NewInt(250),
		Raw: types.Log{
			BlockNumber: 150,
		},
	}

	transferThree := erc20.Erc20Transfer{
		From:  addressB,
		To:    addressZero,
		Value: big.NewInt(250),
		Raw: types.Log{
			BlockNumber: 175,
		},
	}

	transfers := []*erc20.Erc20Transfer{&transferOne, &transferTwo, &transferThree}
	balHistory := GenerateBalanceHistory(&allHolders, &transfers, uint64(startBlock), uint64(endBlock))

	addressAHistory := (*balHistory)[addressA]
	addressBHistory := (*balHistory)[addressB]

	if (*addressAHistory)[0].Uint64() != 0 {
		t.FailNow()
	}

	if (*addressBHistory)[0].Uint64() != 0 {
		t.FailNow()
	}

	if (*addressAHistory)[25].Uint64() != 500 {
		// we should have 500 on block 125
		t.FailNow()
	}

	if (*addressBHistory)[25].Uint64() != 0 {
		t.FailNow()
	}

	if (*addressAHistory)[50].Uint64() != 250 {
		// we should have 250 on block 150
		t.FailNow()
	}

	if (*addressBHistory)[50].Uint64() != 250 {
		// we should have 250 on block 150
		t.FailNow()
	}

	if (*addressAHistory)[99].Uint64() != 250 {
		t.FailNow()
	}

	if (*addressBHistory)[99].Uint64() != 0 {
		t.FailNow()
	}
}

func TestGetAverageBalance(t *testing.T) {
	balances := []*big.Int{big.NewInt(5), big.NewInt(5), big.NewInt(5)}
	result := getAverageBalance(&balances)
	if result.Uint64() != 5 {
		t.FailNow()
	}

	balances = []*big.Int{big.NewInt(10), big.NewInt(5)}
	result = getAverageBalance(&balances)
	if result.Uint64() != 7 {
		t.FailNow()
	}

}
