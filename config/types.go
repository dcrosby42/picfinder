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
	Enabled bool     `yaml:"enabled"`
	Db      DbConfig `yaml:"db"`
}

func (me *EnvConfig) Validate() error {
	if me.Enabled {
		return me.Db.Validate()
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
