package config

import (
	"fmt"

	"github.com/go-yaml/yaml"
	"github.com/urfave/cli"
)

func DumpConfigCommand() cli.Command {
	return cli.Command{
		Name:  "dumpconfig",
		Usage: "Read and dump the config",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			fname := c.GlobalString("config")
			env := c.GlobalString("env")

			fmt.Printf("Config file: %q\n", fname)
			fmt.Printf("Env: %s\n", env)

			config, err := LoadConfig(fname, env, true, false)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}

			err = PrintConfig(config)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil

		},
	}
}

func PrintConfig(config *PicfinderConfig) error {
	serialized, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(serialized))
	return nil
}
