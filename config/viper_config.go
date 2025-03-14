package config

import (
	"github.com/spf13/viper"
)

func ConfigSet() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var configs Config
	if err := viper.Unmarshal(&configs); err != nil {
		return Config{}, err
	}

	return configs, nil
}
