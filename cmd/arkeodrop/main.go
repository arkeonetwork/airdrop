package main

import (
	"math/rand"
	"time"

	"github.com/ArkeoNetwork/airdrop/cli"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cli.Execute()
}
