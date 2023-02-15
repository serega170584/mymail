package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/domain"
	notificator "awesomeProject/internal/proto"
	"google.golang.org/grpc"

	"fmt"
	"log"
	"net"
)

func main() {
	const NetworkLayerTypeTcp = "tcp"

	mainConfig, err := config.NewConfig()
	if err != nil {
		log.Printf("Config handle error: %s", err.Error())
		return
	}

	lis, err := net.Listen(NetworkLayerTypeTcp, fmt.Sprintf("%s:%s", mainConfig.App.Host, mainConfig.App.Port))
	if err != nil {
		log.Printf("failed to listen: %s", err.Error())
		return
	}

	s := grpc.NewServer()
	notificator.RegisterNotificatorServer(s, &domain.NotificatorServer{})
	log.Printf("server listening at %s", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %s", err.Error())
	}

	err = fmt.Errorf("server listening at %s has interrupted", lis.Addr())
	fmt.Println(err.Error())
}
