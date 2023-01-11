package utils

import (
	"math"
	"math/big"
)

func BigIntToFloat(value *big.Int, decimals uint8) float64 {
	transferValue := new(big.Float).SetInt(value)
	transferValue.Quo(transferValue, big.NewFloat(float64(math.Pow10(int(decimals)))))
	transferValueDecimal, _ := transferValue.Float64()
	return transferValueDecimal
}
