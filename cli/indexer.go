package cli

import (
	"github.com/ArkeoNetwork/airdrop/indexer"
	"github.com/ArkeoNetwork/airdrop/pkg/db"
	"github.com/ArkeoNetwork/directory/pkg/config"
	"github.com/ArkeoNetwork/directory/pkg/logging"
	"github.com/spf13/cobra"
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

func runIndexer(cmd *cobra.Command, args []string) {
	log.Info("starting data generation process")
	// flag.Parse()
	c := &Config{}
	flags := cmd.InheritedFlags()
	envPath, _ := flags.GetString("env")
	if envPath == "" {
		if err := config.LoadFromEnv(c, configNames...); err != nil {
			log.Panicf("failed to load config from env: %+v", err)
		}
	} else {
		if err := config.Load(envPath, c); err != nil {
			log.Panicf("failed to load config: %+v", err)
		}
	}

	indexerApp := indexer.NewIndexer(indexer.IndexerAppParams{
		SnapshotStart: c.SnapshotStart,
		SnapshotEnd:   c.SnapshotEnd,
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
