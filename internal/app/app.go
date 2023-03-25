package app

import (
	notificator "awesomeProject/internal/proto"
	"fmt"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"awesomeProject/internal/logger"
	"awesomeProject/internal/server"
)

type App struct {
	config *viper.Viper
	logger *logger.Logger
}

func New(config *viper.Viper) *App {
	appLogger := logger.New(config.GetString("app.env") == "prod", config.GetBool("app.debug"), config.GetString("app.log.filename"))
	return &App{config, appLogger}
}

func (app *App) Run() error {
	const NetworkLayerTypeTcp = "tcp"

	appLogger := app.logger

	config := app.config

	lis, err := net.Listen(NetworkLayerTypeTcp, fmt.Sprintf("%s:%s", config.GetString("app.host"), config.GetString("app.port")))
	if err != nil {
		appLogger.Error(fmt.Sprintf("failed to listen: %s", err.Error()))
		return err
	}

	s := grpc.NewServer()
	notificatorServer := server.New(config, appLogger)
	notificator.RegisterNotificatorServer(s, notificatorServer)
	appLogger.Info(fmt.Sprintf("server listening at %s", lis.Addr()))
	if err := s.Serve(lis); err != nil {
		appLogger.Error(fmt.Sprintf("failed to serve: %s", err.Error()))
		return err
	}

	return nil
}
