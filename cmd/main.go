package main

import (
	"awesomeProject/internal/app"
	"awesomeProject/internal/config"

	"fmt"
	"log"
)

func main() {
	appConfig := config.New()

	appApp := app.New(appConfig)
	err := appApp.Run()

	if err != nil {
		log.Println(fmt.Sprintf("server listening at has interrupted %s\n", err.Error()))
	}
}
