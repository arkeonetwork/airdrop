package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/ArkeoNetwork/airdrop/pkg/db"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type ExportParams struct {
	db.DBConfig
}

var (
	exportCmd = &cobra.Command{
		Use:   "export [chain] [token]",
		Short: "export aggregate data",
		Run:   runExport,
		Args:  cobra.ExactValidArgs(2),
	}
)

func init() {
	exportCmd.Flags().StringP("output", "f", "/tmp/arkeodrop.csv", "csv output file")
}

func runExport(cmd *cobra.Command, args []string) {
	log.Infof("starting export process for %s, %s", args[0], args[1])
	flags := cmd.InheritedFlags()
	envPath, _ := flags.GetString("env")
	c := readConfig(envPath)

	flags = cmd.Flags()
	fileName, _ := flags.GetString("output")
	params := db.DBConfig{
		Host:         c.DBHost,
		Port:         c.DBPort,
		User:         c.DBUser,
		Pass:         c.DBPass,
		DBName:       c.DBName,
		PoolMaxConns: c.DBPoolMaxConns,
		PoolMinConns: c.DBPoolMinConns,
		SSLMode:      c.DBSSLMode,
	}

	err := exportWeightedTokenAvgs(params, args[0], args[1], fileName)
	if err != nil {
		log.Errorf("error exporting: %+v", err)
	}
}

func exportWeightedTokenAvgs(dbConfig db.DBConfig, chain, token, fileName string) error {
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
