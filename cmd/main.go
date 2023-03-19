package main

import (
	"awesomeProject/internal/app"
	"awesomeProject/internal/config"

	"fmt"
	"log"
)

func main() {
	config := config.New()

	app := app.New(config)
	err := app.Run()

	if err != nil {
		log.Println(fmt.Sprintf("server listening at has interrupted %s\n", err.Error()))
	}
}
