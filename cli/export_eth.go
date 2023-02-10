package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/ArkeoNetwork/airdrop/pkg/db"
	"github.com/ArkeoNetwork/common/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type ExportParams struct {
	db.DBConfig
}

var (
	exportCmd = &cobra.Command{
		Use:   "export-eth [chain] [token]",
		Short: "export aggregate data",
		Run:   runExportEthTokenAvg,
		Args:  cobra.ExactValidArgs(2),
	}
)

func init() {
	exportCmd.Flags().StringP("output", "f", "", "csv output file, default /tmp/airdrop_{chain}_{token}.csv")
}

// export block weighted token averages for given eth chain and token
func runExportEthTokenAvg(cmd *cobra.Command, args []string) {
	log.Infof("starting export process for %s, %s", args[0], args[1])
	flags := cmd.InheritedFlags()
	envPath, _ := flags.GetString("env")
	c := utils.ReadDBConfig(envPath)
	if c == nil {
		cmd.PrintErrln("db config undefined")
		return
	}
	flags = cmd.Flags()
	fileName, _ := flags.GetString("output")
	if fileName == "" {
		fileName = fmt.Sprintf("/tmp/airdrop_%s_%s.csv", args[0], args[1])
	}

	err := exportWeightedTokenAvgs(*c, args[0], args[1], fileName)
	if err != nil {
		log.Errorf("error exporting: %+v", err)
	}
}

func exportWeightedTokenAvgs(dbConfig utils.DBConfig, chain, token, fileName string) error {
	d, err := db.New(dbConfig)
	if err != nil {
		return errors.Wrapf(err, "error connecting to the db")
	}

	avgs, err := d.FindAveragedBalances(chain, token)
	if err != nil {
		return errors.Wrapf(err, "error finding averages")
	}
	log.Debugf("found %d averages for %s:%s", len(avgs), chain, token)
	sb := strings.Builder{}
	fmt.Fprint(&sb, "address,balance\n")
	for _, a := range avgs {
		fmt.Fprintf(&sb, "\"%s\",%.18f\n", a.Address, a.Holding)
	}

	if err = os.WriteFile(fileName, []byte(sb.String()), os.ModePerm); err != nil {
		return errors.Wrapf(err, "error writing %s", fileName)
	}
	fmt.Printf("wrote %s", fileName)
	return nil
}
