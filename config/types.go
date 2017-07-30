package config

import "fmt"

type PicfinderConfig struct {
	Envs EnvRootConfig `yaml:"envs"`
}

func (me *PicfinderConfig) Validate() error {
	return me.Envs.Validate()
}

type EnvRootConfig struct {
	Prod    EnvConfig `yaml:"prod"`
	Dev     EnvConfig `yaml:"dev"`
	Test    EnvConfig `yaml:"test"`
	Current EnvConfig `yaml:"-"`
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

func (me EnvRootConfig) ForEnv(envname string) *EnvConfig {
	switch envname {
	case "prod":
		return &me.Prod
	case "dev":
		return &me.Dev
	case "test":
		return &me.Test
	default:
		return &me.Dev
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
