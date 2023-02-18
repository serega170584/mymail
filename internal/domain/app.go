package domain

import (
	notificator "awesomeProject/internal/proto"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	cfg *viper.Viper
}

func NewApp(config *viper.Viper) *App {
	return &App{cfg: config}
}

func (l *App) Run() error {
	const NetworkLayerTypeTcp = "tcp"

	logger := CustomLogger{}

	cfg := l.cfg

	lis, err := net.Listen(NetworkLayerTypeTcp, fmt.Sprintf("%s:%s", cfg.GetString("app.host"), cfg.GetString("app.port")))
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
