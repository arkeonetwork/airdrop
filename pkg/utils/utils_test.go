package utils

import (
	"math/big"
	"testing"

	"gotest.tools/assert"
)

func TestBigIntoFloat(t *testing.T) {
	value := new(big.Int)
	value, ok := value.SetString("100050000000000000000", 10) // 100.05 tokens
	if !ok {
		t.Fatal("failed to set string")
	}
	decimals := uint8(18)
	transferValueDecimal := BigIntToFloat(value, decimals)
	assert.Equal(t, 100.05, transferValueDecimal)
}
