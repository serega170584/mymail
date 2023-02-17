package domain

import (
	"awesomeProject/internal/config"
	notificator "awesomeProject/internal/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	cfg *config.Config
}

func NewApp(config *config.Config) *App {
	return &App{cfg: config}
}

func (l *App) Run() error {
	const NetworkLayerTypeTcp = "tcp"

	logger := CustomLogger{}

	cfg := l.cfg

	lis, err := net.Listen(NetworkLayerTypeTcp, fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port))
	if err != nil {
		logger.Error(fmt.Sprintf("failed to listen: %s", err.Error()))
		return err
	}

	s := grpc.NewServer()
	notificatorServer := NewNotificatorServer()
	notificator.RegisterNotificatorServer(s, notificatorServer)
	logger.Info(fmt.Sprintf("server listening at %s", lis.Addr()))
	if err := s.Serve(lis); err != nil {
		logger.Error(fmt.Sprintf("failed to serve: %s", err.Error()))
		return err
	}

	return nil
}
