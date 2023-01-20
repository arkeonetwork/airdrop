package main

import (
	"flag"

	"github.com/ArkeoNetwork/airdrop/indexer"
	"github.com/ArkeoNetwork/common/logging"
	"github.com/ArkeoNetwork/common/utils"
)

type Config struct {
	SnapshotStart  uint64 `mapstructure:"SNAPSHOT_START"`
	SnapshotEnd    uint64 `mapstructure:"SNAPSHOT_END"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         uint   `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
	DBSSLMode      string `mapstructure:"DB_SSL_MODE"`
	DBPoolMaxConns int    `mapstructure:"DB_POOL_MAX_CONNS"`
	DBPoolMinConns int    `mapstructure:"DB_POOL_MIN_CONNS"`
}

var (
	log         = logging.WithoutFields()
	envPath     = flag.String("env", "", "path to env file (default: use os env)")
	configNames = []string{
		"SNAPSHOT_START",
		"SNAPSHOT_END",
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
		if err := utils.LoadFromEnv(c, configNames...); err != nil {
			log.Panicf("failed to load config from env: %+v", err)
		}
	} else {
		if err := utils.Load(*envPath, c); err != nil {
			log.Panicf("failed to load config: %+v", err)
		}
	}

	dbc := utils.ReadDBConfig(*envPath)
	if dbc == nil {
		log.Error("db config undefined")
		return
	}
	indexerApp := indexer.NewIndexer(indexer.IndexerAppParams{
		SnapshotStart: c.SnapshotStart,
		SnapshotEnd:   c.SnapshotEnd,
		DBConfig:      *dbc,
	})
	indexerApp.Start()
	log.Debug("finished data generation process")
}
