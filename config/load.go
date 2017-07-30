package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/urfave/cli"
)

func GetConfig(c *cli.Context) (*PicfinderConfig, error) {
	fname := c.GlobalString("config")
	env := c.GlobalString("env")
	return LoadConfig(fname, env, true, true)
}

func LoadConfig(fname string, env string, initDefaultFile bool, setCurrent bool) (*PicfinderConfig, error) {
	_, err := os.Stat(fname)
	if err != nil {
		if os.IsNotExist(err) && initDefaultFile && fname == DefaultConfigFile {
			var err2 error
			err2 = GenerateDefaultConfigFile()
			if err2 != nil {
				return nil, err2
			}
		} else {
			return nil, err
		}
	}
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

	if setCurrent {
		envCfg, err := config.Envs.ForEnv(env)
		if err != nil {
			return nil, err
		}
		if !envCfg.Enabled {
			return nil, fmt.Errorf("Environment %q is not enabled", env)
		}
		config.Envs.Current = *envCfg
		config.Envs.CurrentName = env
	}

	return config, nil
}

func GenerateDefaultConfigFile() error {
	example := ExampleConfig()

	data, err := yaml.Marshal(example)
	if err != nil {
		return err
	}

	fmt.Printf("Generating default configuration in file %q\n", DefaultConfigFile)
	err = ioutil.WriteFile(DefaultConfigFile, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
