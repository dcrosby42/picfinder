package main

import (
	"os"

	"github.com/dcrosby42/picfinder/commands/db"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, commands_db.Command())

	app.Run(os.Args)
}
