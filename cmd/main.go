package main

import (
	"log"

	"awesomeProject/internal/app"
	"awesomeProject/internal/config"
)

func main() {
	appConfig := config.New()

	appApp := app.New(appConfig)
	err := appApp.Run()

	if err != nil {
		log.Printf("server listening at has interrupted %s\n", err.Error())
	}
}
