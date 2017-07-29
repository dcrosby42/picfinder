package main

import (
	"os"

	"github.com/dcrosby42/picfinder/dbutil"
	"github.com/dcrosby42/picfinder/sandbox"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		dbutil.Command(),
		sandbox.Command(),
	}

	app.Run(os.Args)
}
