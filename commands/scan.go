package commands

import (
	"github.com/dcrosby42/picfinder/api_client"
	"github.com/dcrosby42/picfinder/scan"
	"github.com/urfave/cli"
)

func ScanCommand() cli.Command {
	return cli.Command{
		Name:  "scan",
		Usage: "File scanning operations",
		Subcommands: []cli.Command{
			scan_info_command(),
			scan_update_command(),
		},
	}
}

func scanFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "dir",
			Usage: "The dir to scan",
			Value: "/Users/crosby/Pictures",
		},
		cli.BoolFlag{
			Name:  "all",
			Usage: "Scan all the files, not just media files",
		},
	}
}

func scan_info_command() cli.Command {
	return cli.Command{
		Name:  "info",
		Usage: "Print media file count-ups for a directory.",
		Flags: scanFlags(),
		Action: func(c *cli.Context) error {
			err := scan.PrintFileTypeSummary("localhost", c.String("dir"), c.Bool("all"))
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil
		},
	}
}

func scan_update_command() cli.Command {
	return cli.Command{
		Name:  "update",
		Usage: "Scan files on local drive and send updates to the picfinder server API.",
		Flags: append(scanFlags(), api_client.RemoteServerFlags()...),
		Action: func(c *cli.Context) error {
			host := c.String("host")
			port := c.String("port")
			client, closeConn, err := api_client.NewClient_HostPort(host, port)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}

			defer closeConn()

			err = scan.ScanAndSend(client, "xbb", c.String("dir"), c.Bool("all"))
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil
		},
	}
}
