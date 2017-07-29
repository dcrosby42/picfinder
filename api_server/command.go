package api_server

import (
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "Start the Picfinder GRPC API Server",
		// Subcommands: []cli.Command{
		// },
		Action: func(c *cli.Context) error {
			err := BuildAndListen(":13131")
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil

		},
	}
}
