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
		Use:   "index",
		Short: "gather chain data store in our db",
		Run:   runIndexer,
	}
)

func init() {
	rootCmd.LocalFlags().StringP("env", "e", "docker/dev/docker.env", "env file to source")
	rootCmd.InheritedFlags().StringP("env", "e", "docker/dev/docker.env", "env file to source")
	rootCmd.PersistentFlags().StringP("env", "e", "docker/dev/docker.env", "env file to source")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(indexCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
