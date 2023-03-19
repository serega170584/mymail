package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func New() *viper.Viper {
	config := viper.New()
	config.SetConfigFile("./config/local.json")

	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		fmt.Printf("fatal error config file: %s", err.Error())
	}

	return config
}
