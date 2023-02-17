package config

import (
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
	config := &viper.Viper{}

	config.SetConfigName("local")
	config.SetConfigType("json")
	config.AddConfigPath("./config")

	return config
}
