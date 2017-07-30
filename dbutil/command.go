package dbutil

import (
	"fmt"

	"github.com/dcrosby42/picfinder/commands"
	"github.com/dcrosby42/picfinder/config"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:  "db",
		Usage: "Tools for manipulating the database",
		Subcommands: []cli.Command{
			db_rebuild_command(),
			db_shellenv_command(),
		},
	}
}

func db_rebuild_command() cli.Command {
	return commands.ProtectedCommand(cli.Command{
		Name:  "rebuild",
		Usage: "Rebuild the entire db from scratch. WILL DESTROY ALL DATA",
		Action: func(c *cli.Context) error {
			cfg, err := config.GetConfig(c)
			if err != nil {
				return cli.NewExitError(err.Error(), -42)
			}

			db, err := ConnectDatabase(cfg.Envs.Current.Server.Db)
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

func db_shellenv_command() cli.Command {
	return cli.Command{
		Name:  "shellvars",
		Usage: "Print out shell env vars based on current config/env",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			fname := c.GlobalString("config")
			env := c.GlobalString("env")

			// fmt.Printf("# Config file: %q, Env: %q:\n", fname, env)

			config, err := config.LoadConfig(fname, env, true, true)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}

			dbConf := config.Envs.Current.Server.Db
			fmt.Printf("export PICFINDER_ENV=%s\n", config.Envs.CurrentName)
			fmt.Printf("export PICFINDER_CONFIG=%s\n", fname)
			fmt.Printf("export DBNAME=%s\n", dbConf.DbName)
			fmt.Printf("export DBUSER=%s\n", dbConf.Username)
			fmt.Printf("export DBPASS=%s\n", dbConf.Password)
			fmt.Printf("export DBHOST=%s\n", dbConf.Host)
			fmt.Printf("export DBPORT=%d\n", dbConf.Port)
			return nil

		},
	}
}
