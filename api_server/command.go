package api_server

import (
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
			err := BuildAndListen(c.String("bind"))
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil

		},
	}
}
