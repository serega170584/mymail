package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/domain"
	"fmt"
)

func main() {
	logger := &domain.CustomLogger{}

	mainConfig := config.NewConfig()

	app := domain.NewApp(mainConfig)
	err := app.Run()

	if err != nil {
		logger.Error(fmt.Sprintf("server listening at has interrupted %s\n", err.Error()))
	}
}
