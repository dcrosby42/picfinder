package config

import "github.com/urfave/cli"

func GlobalFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Usage: "the config file",
			Value: "local-config.yaml",
		},
		cli.StringFlag{
			Name:  "env",
			Usage: "the environment name to use in config",
			Value: "dev",
		},
	}
}
