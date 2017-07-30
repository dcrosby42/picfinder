package config

func ExampleConfig() *PicfinderConfig {
	cfg := &PicfinderConfig{}
	cfg.Envs.Prod = EnvConfig{
		Enabled: false,
		Db: DbConfig{
			DbName:   "picfinder",
			Username: "picfinder",
			Password: "picfinder",
			Host:     "localhost",
			Port:     3306,
		},
	}
	cfg.Envs.Dev = EnvConfig{
		Enabled: true,
		Db: DbConfig{
			DbName:   "picfinder_dev",
			Username: "picfinder_dev",
			Password: "picfinder_dev",
			Host:     "localhost",
			Port:     3306,
		},
	}
	cfg.Envs.Test = EnvConfig{
		Enabled: false,
		Db: DbConfig{
			DbName:   "picfinder_test",
			Username: "picfinder_test",
			Password: "picfinder_test",
			Host:     "localhost",
			Port:     3306,
		},
	}
	return cfg

}
