package api_server

import (
	"github.com/dcrosby42/picfinder/config"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "Start the Picfinder GRPC API Server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "bind",
				Usage: "Server bind addr",
				Value: ":13131",
			},
		},
		Action: func(c *cli.Context) error {
			cfg, err := config.GetConfig(c)
			if err != nil {
				return err
			}
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}

			err = BuildAndListen(c.String("bind"), cfg.Envs.Current.Db)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil

		},
	}
}
