package cli

import (
	"fmt"

	"github.com/ArkeoNetwork/airdrop/indexer"
	"github.com/ArkeoNetwork/common/utils"
	"github.com/spf13/cobra"
)

var (
	indexDelegatorsCmd = &cobra.Command{
		Use:   "delegators",
		Short: "gather cosmos-sdk chain data store in our db",
		Run:   runDelegatorsIndexer,
		Args:  cobra.ExactValidArgs(1),
	}
)

func runDelegatorsIndexer(cmd *cobra.Command, args []string) {
	flags := cmd.InheritedFlags()
	envPath, _ := flags.GetString("env")
	c := utils.ReadDBConfig(envPath)
	if c == nil {
		fmt.Print("error: no config loaded")
		return
	}

	chain := args[0]
	params := indexer.CosmosIndexerParams{Chain: chain, DB: *c}
	indxr, err := indexer.NewCosmosIndexer(params)
	if err != nil {
		cmd.PrintErrf("error creating cosmos indexer: %+v", err)
		return
	}

	if err := indxr.IndexDelegators(); err != nil {
		fmt.Printf("error indexing delegators: %+v", err)
		return
	}
}
