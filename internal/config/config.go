package config

import (
	"github.com/spf13/viper"
)

func Load() error {
	viper.SetConfigFile("../.env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.AutomaticEnv()

	return nil
}
