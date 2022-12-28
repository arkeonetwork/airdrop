package main

import (
	"flag"

	"github.com/ArkeoNetwork/directory/pkg/config"
	"github.com/ArkeoNetwork/directory/pkg/logging"
)

type Config struct {
	EthRPC           string `mapstructure:"ETH_RPC"`
	FoxAddressEth    string `mapstructure:"FOX_ADDRESS_ETH"`
	FoxAddressGnosis string `mapstructure:"FOX_ADDRESS_GNOSIS"`
}

var (
	log         = logging.WithoutFields()
	envPath     = flag.String("env", "", "path to env file (default: use os env)")
	configNames = []string{
		"ETH_RPC",
		"FOX_ADDRESS_ETH",
		"FOX_ADDRESS_GNOSIS",
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
}
