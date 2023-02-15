package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	App struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	Mail struct {
		Port     int    `json:"port"`
		From     string `json:"from"`
		To       string `json:"to"`
		Host     string `json:"host"`
		Password string `json:"password"`
	}
}

func NewConfig() (*Config, error) {
	config := &Config{}

	file, err := os.Open("./config/local.json")
	if err != nil {
		log.Printf("failed to open config: %s", err.Error())
		return &Config{}, err
	}
	defer file.Close()

	d := json.NewDecoder(file)

	if err := d.Decode(config); err != nil {
		log.Printf("failed to decode config: %s", err.Error())
		return &Config{}, err
	}

	return config, nil
}
