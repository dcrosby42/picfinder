package config

import "fmt"

type PicfinderConfig struct {
	Envs EnvRootConfig `yaml:"envs"`
}

func (me *PicfinderConfig) Validate() error {
	return me.Envs.Validate()
}

type EnvRootConfig struct {
	Prod        EnvConfig `yaml:"prod"`
	Dev         EnvConfig `yaml:"dev"`
	Test        EnvConfig `yaml:"test"`
	Current     EnvConfig `yaml:"-"`
	CurrentName string    `yaml:"-"`
}

func (me *EnvRootConfig) Validate() error {
	var err error
	err = me.Prod.Validate()
	if err != nil {
		return err
	}
	err = me.Dev.Validate()
	if err != nil {
		return err
	}
	err = me.Dev.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (me EnvRootConfig) ForEnv(envname string) (*EnvConfig, error) {
	switch envname {
	case "prod":
		return &me.Prod, nil
	case "dev":
		return &me.Dev, nil
	case "test":
		return &me.Test, nil
	default:
		return nil, fmt.Errorf("Invalid env name %q", envname)
	}
}

type EnvConfig struct {
	Enabled bool         `yaml:"enabled"`
	Server  ServerConfig `yaml:"server"`
	Client  ClientConfig `yaml:"client"`
}

func (me *EnvConfig) Validate() error {
	if me.Enabled {
		var err error
		err = me.Server.Validate()
		if err != nil {
			return err
		}
		err = me.Client.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

type ServerConfig struct {
	Enabled   bool            `yaml:"enabled"`
	Db        DbConfig        `yaml:"db"`
	ApiServer ApiServerConfig `yaml:"api_client"`
}

func (me *ServerConfig) Validate() error {
	if me.Enabled {
		var err error
		err = me.Db.Validate()
		if err != nil {
			return err
		}
		err = me.ApiServer.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

type ClientConfig struct {
	Enabled   bool            `yaml:"enabled"`
	ApiClient ApiClientConfig `yaml:"api_client"`
}

func (me *ClientConfig) Validate() error {
	if me.Enabled {
		var err error
		err = me.ApiClient.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

type DbConfig struct {
	DbName   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

func (me *DbConfig) Validate() error {
	if me.DbName == "" {
		return fmt.Errorf("dbname must not be blank")
	}
	if me.Username == "" {
		return fmt.Errorf("username must not be blank")
	}
	if me.Password == "" {
		return fmt.Errorf("password must not be blank")
	}
	return nil
}

type ApiClientConfig struct {
	Server string `yaml:"server"`
	Port   string `yaml:"port"`
}

func (me *ApiClientConfig) Validate() error {
	if me.Server == "" {
		return fmt.Errorf("api_client.server must not be blank")
	}
	if me.Port == "" {
		return fmt.Errorf("api_client.port must not be blank")
	}
	return nil
}

type ApiServerConfig struct {
	BindAddr string `yaml:"bind_addr"`
}

func (me *ApiServerConfig) Validate() error {
	if me.BindAddr == "" {
		return fmt.Errorf("api_server.bind_addr must not be blank")
	}
	return nil
}
