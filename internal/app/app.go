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
	logger := logger.New(config.GetString("app.env") == "prod", config.GetBool("app.debug"), config.GetString("app.log.filename"), config.GetString("app.log.prefix"))
	return &App{config, logger}
}

func (app *App) Run() error {
	const NetworkLayerTypeTcp = "tcp"

	logger := app.logger

	lis, err := net.Listen(NetworkLayerTypeTcp, fmt.Sprintf("%s:%s", app.config.GetString("app.host"), app.config.GetString("app.port")))
	if err != nil {
		logger.Error(fmt.Sprintf("failed to listen: %s", err.Error()))
		return err
	}

	s := grpc.NewServer()
	notificatorServer := server.New(app.config, app.logger)
	notificator.RegisterNotificatorServer(s, notificatorServer)
	logger.Info(fmt.Sprintf("server listening at %s", lis.Addr()))
	if err := s.Serve(lis); err != nil {
		logger.Error(fmt.Sprintf("failed to serve: %s", err.Error()))
		return err
	}

	return nil
}
