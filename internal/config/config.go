package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	defaultConfigPath = "./config/local.yaml"
)

type Config struct {
	Env               string `yaml:"env" env-required:"true"`
	HTTPServerAddress string `yaml:"http_server_address" env-required:"true"`
}

// MustLoad function loads config from configPath
func MustLoad() *Config {

	// get path of configuration file
	configPath := defaultConfigPath

	// check if config file exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file doesn't exist:" + configPath)
	}

	// read config
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		panic("unable to read config file")
	}

	return &cfg
}
