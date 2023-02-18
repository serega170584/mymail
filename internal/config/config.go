package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Env   string `json:"env"`
		Debug bool   `json:"debug"`
		Host  string `json:"host"`
		Port  string `json:"port"`
	}
	Mail struct {
		Port int    `json:"port"`
		From string `json:"from"`
		To   string `json:"to"`
		Host string `json:"host"`
	}
}

func NewConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigFile("./config/local.json")

	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		fmt.Printf("fatal error config file: %s", err.Error())
	}

	return config
}
