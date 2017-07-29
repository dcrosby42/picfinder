package commands

import "github.com/urfave/cli"

func ProtectedCommand(cmd cli.Command) cli.Command {
	if cmd.Flags == nil {
		cmd.Flags = []cli.Flag{}
	}
	cmd.Flags = append(cmd.Flags, ReallyFlag())

	actualAction := cmd.Action.(func(*cli.Context) error)
	cmd.Action = func(c *cli.Context) error {
		if !c.Bool("really") {
			return cli.NewExitError("You need to be serious; try using -really.", -42)
		}
		return actualAction(c)
	}
	return cmd
}

func ReallyFlag() cli.Flag {
	return cli.BoolFlag{
		Name:  "really",
		Usage: "Actually do this",
	}
}
