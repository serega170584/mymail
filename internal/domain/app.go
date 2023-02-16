package domain

import (
	"awesomeProject/internal/config"
	notificator "awesomeProject/internal/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
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

	cfg := l.cfg

	lis, err := net.Listen(NetworkLayerTypeTcp, fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port))
	if err != nil {
		log.Printf("failed to listen: %s", err.Error())
		return err
	}

	s := grpc.NewServer()
	notificatorServer := NewNotificatorServer()
	notificator.RegisterNotificatorServer(s, notificatorServer)
	log.Printf("server listening at %s", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %s", err.Error())
		return err
	}

	return nil
}
