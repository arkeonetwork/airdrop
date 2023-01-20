package cli

import (
	"github.com/ArkeoNetwork/airdrop/indexer"
	"github.com/ArkeoNetwork/common/logging"
	"github.com/ArkeoNetwork/common/utils"
	"github.com/spf13/cobra"
)

type Config struct {
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
	flags := cmd.InheritedFlags()
	envPath, _ := flags.GetString("env")
	c := utils.ReadDBConfig(envPath)
	if c == nil {
		cmd.PrintErr("db config undefined")
		return
	}

	indexerApp := indexer.NewIndexer(indexer.IndexerAppParams{
		DBConfig: *c,
	})

	indexerApp.Start()
	log.Debug("finished data generation process")
}
