package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		Database       DatabaseConfig
		Authentication Authentication
		Server         ServerConfig
	}

	DatabaseConfig struct {
		Host     string
		Port     int
		Name     string
		User     string
		Password string
	}

	Authentication struct {
		Key string
	}

	ServerConfig struct {
		Port     int
		Debug    bool
		TimeZone string
	}
)

func Load(nome string, path ...string) (config *Config, err error) {

	viper.SetConfigName(nome)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	for _, p := range path {
		viper.AddConfigPath(p)
	}
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	config = &Config{}
	err = viper.Unmarshal(config)

	return
}
