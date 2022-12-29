package main

import (
	"flag"

	"github.com/ArkeoNetwork/directory/pkg/config"
	"github.com/ArkeoNetwork/directory/pkg/logging"
	"github.com/ArkeoNetwork/merkle-drop/datagen"
)

type Config struct {
	EthRPC           string `mapstructure:"ETH_RPC"`
	FoxGenesisBlock  uint64 `mapstructure:"FOX_GENESIS_BLOCK"`
	FoxAddressEth    string `mapstructure:"FOX_ADDRESS_ETH"`
	FoxAddressGnosis string `mapstructure:"FOX_ADDRESS_GNOSIS"`
	SnapshotStart    uint64 `mapstructure:"SNAPSHOT_START"`
	SnapshotEnd      uint64 `mapstructure:"SNAPSHOT_END"`
}

var (
	log         = logging.WithoutFields()
	envPath     = flag.String("env", "", "path to env file (default: use os env)")
	configNames = []string{
		"ETH_RPC",
		"FOX_GENESIS_BLOCK",
		"FOX_ADDRESS_ETH",
		"FOX_ADDRESS_GNOSIS",
		"SNAPSHOT_START",
		"SNAPSHOT_END",
	}
)

func main() {
	log.Info("starting data generation process")
	flag.Parse()
	c := &Config{}
	if *envPath == "" {
		if err := config.LoadFromEnv(c, configNames...); err != nil {
			log.Panicf("failed to load config from env: %+v", err)
		}
	} else {
		if err := config.Load(*envPath, c); err != nil {
			log.Panicf("failed to load config: %+v", err)
		}
	}

	datagen.NewApp(datagen.AppParams{
		EthRPC:           c.EthRPC,
		FoxGenesisBlock:  c.FoxGenesisBlock,
		FoxAddressEth:    c.FoxAddressEth,
		FoxAddressGnosis: c.FoxAddressGnosis,
		SnapshotStart:    c.SnapshotStart,
		SnapshotEnd:      c.SnapshotEnd,
	})
}
