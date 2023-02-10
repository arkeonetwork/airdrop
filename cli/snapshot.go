package cli

import (
	"github.com/ArkeoNetwork/airdrop/pkg/db"
	"github.com/ArkeoNetwork/airdrop/snapshot"
	"github.com/spf13/cobra"
)

func runSnapshotIndexer(cmd *cobra.Command, args []string) {
	log.Info("starting gethering snapshot data process")
	flags := cmd.InheritedFlags()
	envPath, _ := flags.GetString("env")
	c := readConfig(envPath)
	snapshotIndexerApp := snapshot.NewSnapshotIndexer(snapshot.SnapshotIndexerAppParams{
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
	snapshotIndexerApp.Start()
	log.Debug("finished gethering snapshot data process")
}
