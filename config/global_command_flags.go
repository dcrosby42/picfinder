package config

import "github.com/urfave/cli"

const DefaultConfigFile = "local-config.yaml"
const DefaultConfigEnv = "dev"

func GlobalFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Usage: "the config file",
			Value: DefaultConfigFile,
		},
		cli.StringFlag{
			Name:  "env",
			Usage: "the environment name to use in config",
			Value: DefaultConfigEnv,
		},
	}
}

func RemoteServerFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "server",
			Usage: "The api server host",
			Value: "127.0.0.1",
		},
		cli.StringFlag{
			Name:  "port",
			Usage: "The api server port",
			Value: "13131",
		},
	}
}
