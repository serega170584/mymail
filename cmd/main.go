package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/domain"
	"fmt"
	"log"
)

func main() {
	mainConfig, err := config.NewConfig()
	if err != nil {
		log.Printf("Config handle error: %s", err.Error())
		return
	}

	app := domain.NewApp(mainConfig)
	err = app.Run()

	if err != nil {
		fmt.Println("server listening at has interrupted\n", err)
	}
}
