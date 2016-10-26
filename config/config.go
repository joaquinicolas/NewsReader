package config

import ("github.com/BurntSushi/toml"
)

type Config struct {
	RemoteServer string
	Port int
}

func ReadConfig() *Config  {
	var config Config
	if _,err := toml.Decode("../config.toml",&config); err != nil {
		return &config
	}

	return nil
}
