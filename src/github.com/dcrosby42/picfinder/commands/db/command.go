package commands_db

import (
	"fmt"

	"github.com/dcrosby42/picfinder/commands"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:  "db",
		Usage: "Tools for manipulating the database",
		Subcommands: []cli.Command{
			db_create_command(),
		},
	}
}

func db_create_command() cli.Command {
	return commands.ProtectedCommand(cli.Command{
		Name:  "create",
		Usage: "Rebuild the entire db from scratch. WILL DESTROY ALL DATA",
		Action: func(c *cli.Context) error {
			fmt.Println("TBD!")
			return nil
		},
	})
}
