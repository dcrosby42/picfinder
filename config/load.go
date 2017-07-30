package config

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

func LoadConfig(fname string, env string, setCurrent bool) (*PicfinderConfig, error) {
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
