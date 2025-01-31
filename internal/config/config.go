package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigFile("../.env") 
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	viper.AutomaticEnv()

	fmt.Println("config working")
}


