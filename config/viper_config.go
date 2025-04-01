package config

import (
	"SysAuthenticate/logger"
	"github.com/spf13/viper"
)

func ConfigSet() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Errorf("Error reading config file, %s", err)
		return Config{}, err
	}

	var configs Config
	if err := viper.Unmarshal(&configs); err != nil {
		return Config{}, err
	}

	return configs, nil
}
