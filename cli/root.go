package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "arkeodrop",
		Short: "arkeodrop is airdrop utilities",
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "show version",
		Run: func(cmd *cobra.Command, args []string) {
			println("arkeodrop v0.0.1\n")
		},
	}

	indexCmd = &cobra.Command{
		Use:   "index-eth",
		Short: "gather eth chain data store in our db",
		Run:   runIndexer,
	}

	indexCosmosCmd = &cobra.Command{
		Use:   "index-cosmos",
		Short: "cosmos-sdk indexing",
	}
)

func init() {
	rootCmd.PersistentFlags().StringP("env", "e", "docker/dev/docker.env", "env file to source")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(indexCmd)
	indexCosmosCmd.AddCommand(indexDelegatorsCmd)
	indexCosmosCmd.AddCommand(indexThorchainLPCmd)
	indexCosmosCmd.AddCommand(indexOsmoLPCmd)
	indexCosmosCmd.AddCommand(indexStartingBalancesCmd)
	rootCmd.AddCommand(indexCosmosCmd)
	rootCmd.AddCommand(exportCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
