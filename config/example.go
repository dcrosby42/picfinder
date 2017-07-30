package config

func ExampleConfig() *PicfinderConfig {
	cfg := &PicfinderConfig{}
	cfg.Envs.Prod = ExampleEnvConfig("picfinder")
	cfg.Envs.Dev = ExampleEnvConfig("picfinder_dev")
	cfg.Envs.Dev.Enabled = true
	cfg.Envs.Dev.Server.Enabled = true
	cfg.Envs.Dev.Client.Enabled = true
	cfg.Envs.Test = ExampleEnvConfig("picfinder_test")
	return cfg
}
func ExampleEnvConfig(magicWord string) EnvConfig {
	return EnvConfig{
		Enabled: false,
		Server: ServerConfig{
			Db: DbConfig{
				DbName:   magicWord,
				Username: magicWord,
				Password: magicWord,
				Host:     "localhost",
				Port:     3306,
			},
			ApiServer: ApiServerConfig{
				BindAddr: ":13131",
			},
		},
		Client: ClientConfig{
			ApiClient: ExampleApiClientConfig(),
		},
	}
}

func ExampleApiClientConfig() ApiClientConfig {
	return ApiClientConfig{
		Server: "localhost",
		Port:   "13131",
	}
}
