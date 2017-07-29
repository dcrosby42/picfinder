package dbutil

import (
	"github.com/dcrosby42/picfinder/commands"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:  "db",
		Usage: "Tools for manipulating the database",
		Subcommands: []cli.Command{
			db_rebuild_command(),
		},
	}
}

func db_rebuild_command() cli.Command {
	return commands.ProtectedCommand(cli.Command{
		Name:  "rebuild",
		Usage: "Rebuild the entire db from scratch. WILL DESTROY ALL DATA",
		Action: func(c *cli.Context) error {
			db, err := ConnectDatabase()
			if err != nil {
				return cli.NewExitError(err.Error(), -37)
			}
			err = RebuildTables(db)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil
		},
	})
}
