package commands

import (
	"github.com/dcrosby42/picfinder/scan"
	"github.com/urfave/cli"
)

func ScanCommand() cli.Command {
	return cli.Command{
		Name:  "scan",
		Usage: "File scanning operations",
		Subcommands: []cli.Command{
			scan_info_command(),
		},
	}
}

func scan_info_command() cli.Command {
	return cli.Command{
		Name:  "info",
		Usage: "Print media file count-ups for a directory.",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "dir",
				Usage: "The dir to scan",
				Value: "/Users/crosby/Pictures",
			},
			cli.BoolFlag{
				Name:  "all",
				Usage: "Scan all the files, not just media files",
			},
		},
		Action: func(c *cli.Context) error {
			err := scan.PrintFileTypeSummary("localhost", c.String("dir"), c.Bool("media"))
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil
		},
	}
}
