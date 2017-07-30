package config

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
	"github.com/urfave/cli"
)

func RemoteServerFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "server",
			Usage: "The api server host",
			Value: "127.0.0.1",
		},
		cli.StringFlag{
			Name:  "port",
			Usage: "The api server port",
			Value: "13131",
		},
	}
}

func DumpConfigCommand() cli.Command {
	return cli.Command{
		Name:  "dumpconfig",
		Usage: "Read and dump the config",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			fname := "default-config.yaml"

			// Load config
			fmt.Printf("Loading config from %q\n", fname)
			config, err := LoadConfigFile(fname)
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

func LoadConfigFile(fname string) (*PicfinderConfig, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	config := &PicfinderConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	err = config.Validate()
	if err != nil {
		return nil, err
	}
	return config, nil
}
func PrintConfig(config *PicfinderConfig) error {
	serialized, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(serialized))
	return nil
}
