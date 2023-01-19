package cli

import "github.com/ArkeoNetwork/directory/pkg/config"

func readConfig(envPath string) *Config {
	c := &Config{}
	if envPath == "" {
		if err := config.LoadFromEnv(c, configNames...); err != nil {
			log.Panicf("failed to load config from env: %+v", err)
		}
	} else {
		if err := config.Load(envPath, c); err != nil {
			log.Panicf("failed to load config: %+v", err)
		}
	}
	return c
}
