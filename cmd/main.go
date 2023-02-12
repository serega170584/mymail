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

	mainConfig, err := config.NewConfig()
	if err != nil {
		log.Printf("Config handle error: %v", err)
		return
	}

	// TODO: заменить на константу
	// http.StatusOK
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", mainConfig.App.Host, mainConfig.App.Port))
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	notificator.RegisterNotificatorServer(s, &domain.NotificatorServer{})
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}

	// TODO: errorf
	log.Printf("server listening at %s", lis.Addr())
}
