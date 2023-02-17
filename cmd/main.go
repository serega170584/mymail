package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/domain"
	"fmt"
)

func main() {
	logger := &domain.CustomLogger{}

	mainConfig, err := config.NewConfig()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	app := domain.NewApp(mainConfig)
	err = app.Run()

	if err != nil {
		fmt.Println("server listening at has interrupted\n", err)
	}
}
