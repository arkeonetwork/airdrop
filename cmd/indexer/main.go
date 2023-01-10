package main

import (
	"flag"

	"github.com/ArkeoNetwork/directory/pkg/config"
	"github.com/ArkeoNetwork/directory/pkg/logging"
	"github.com/ArkeoNetwork/merkle-drop/indexer"
	"github.com/ArkeoNetwork/merkle-drop/pkg/db"
)

type Config struct {
	EthRPC                 string `mapstructure:"ETH_RPC"`
	FoxGenesisBlock        uint64 `mapstructure:"FOX_GENESIS_BLOCK"`
	FoxLPGenesisBlock      uint64 `mapstructure:"FOX_LP_GENESIS_BLOCK"`
	FoxStakingGenesisBlock uint64 `mapstructure:"FOX_STAKING_GENESIS_BLOCK"`
	FoxStakingAddressEth   string `mapstructure:"FOX_STAKING_ADDRESS_ETH"`
	FoxAddressEth          string `mapstructure:"FOX_ADDRESS_ETH"`
	FoxAddressGnosis       string `mapstructure:"FOX_ADDRESS_GNOSIS"`
	FoxLPAddressEth        string `mapstructure:"FOX_LP_ADDRESS_ETH"`
	SnapshotStart          uint64 `mapstructure:"SNAPSHOT_START"`
	SnapshotEnd            uint64 `mapstructure:"SNAPSHOT_END"`
	SnapshotStartBlockEth  uint64 `mapstructure:"SNAPSHOT_START_BLOCK_ETH"`
	SnapshotEndBlockEth    uint64 `mapstructure:"SNAPSHOT_END_BLOCK_ETH"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 uint   `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	DBSSLMode              string `mapstructure:"DB_SSL_MODE"`
	DBPoolMaxConns         int    `mapstructure:"DB_POOL_MAX_CONNS"`
	DBPoolMinConns         int    `mapstructure:"DB_POOL_MIN_CONNS"`
}

var (
	log         = logging.WithoutFields()
	envPath     = flag.String("env", "", "path to env file (default: use os env)")
	configNames = []string{
		"ETH_RPC",
		"FOX_GENESIS_BLOCK",
		"FOX_LP_GENESIS_BLOCK",
		"FOX_ADDRESS_ETH",
		"FOX_ADDRESS_GNOSIS",
		"FOX_LP_ADDRESS_ETH",
		"FOX_STAKING_ADDRESS_ETH",
		"FOX_STAKING_GENESIS_BLOCK",
		"SNAPSHOT_START",
		"SNAPSHOT_END",
		"SNAPSHOT_START_BLOCK_ETH",
		"SNAPSHOT_END_BLOCK_ETH",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASS",
		"DB_NAME",
		"DB_SSL_MODE",
		"DB_POOL_MAX_CONNS",
		"DB_POOL_MIN_CONNS",
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

	indexerApp := indexer.NewIndexer(indexer.IndexerAppParams{
		EthRPC:                 c.EthRPC,
		FoxGenesisBlock:        c.FoxGenesisBlock,
		FoxLPGenesisBlock:      c.FoxLPGenesisBlock,
		FoxAddressEth:          c.FoxAddressEth,
		FoxAddressGnosis:       c.FoxAddressGnosis,
		FoxLPAddressEth:        c.FoxLPAddressEth,
		FoxStakingGenesisBlock: c.FoxStakingGenesisBlock,
		FoxStakingAddressEth:   c.FoxStakingAddressEth,
		SnapshotStart:          c.SnapshotStart,
		SnapshotEnd:            c.SnapshotEnd,
		SnapshotStartBlockEth:  c.SnapshotStartBlockEth,
		SnapshotEndBlockEth:    c.SnapshotEndBlockEth,
		DBConfig: db.DBConfig{
			Host:         c.DBHost,
			Port:         c.DBPort,
			User:         c.DBUser,
			Pass:         c.DBPass,
			DBName:       c.DBName,
			PoolMaxConns: c.DBPoolMaxConns,
			PoolMinConns: c.DBPoolMinConns,
			SSLMode:      c.DBSSLMode,
		},
	})
	indexerApp.Start()
	log.Debug("finished data generation process")
}
