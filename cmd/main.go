package main

import (
	"awesomeProject/internal/config"
	notificator "awesomeProject/internal/proto"
	"context"
	"crypto/tls"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/gomail.v2"
	"log"
	"net"
	"time"
)

type notificatorServer struct {
	cfg config.Config

	notificator.UnimplementedNotificatorServer
}

// MyMail TODO: сделать ее локальной
func (r *notificatorServer) MyMail(ctx context.Context, in *notificator.EmailRequest) (*emptypb.Empty, error) {
	log.Println("Received:", in.GetTo(), in.GetSubject())

	message := gomail.NewMessage()

	// tood: default in config
	message.SetHeader("From", mainConfig.Mail.From)
	// todo: take from request
	message.SetHeader("To", mainConfig.Mail.To)
	message.SetHeader("Subject", fmt.Sprintf("grpc handler was triggered at %s", time.Now().String()))

	// TODO: google mailchimp если сложно то найдем другое решение
	dialer := gomail.NewDialer(mainConfig.Mail.Host, mainConfig.Mail.Port, mainConfig.Mail.From, mainConfig.Mail.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		log.Printf("failed to send mail: %v\n", err)
		log.Printf("failed to send mail: %s\n", err.Error())
		return &emptypb.Empty{}, err
	}
	log.Printf("Letter is sent")

	return &emptypb.Empty{}, nil
}

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
	notificator.RegisterNotificatorServer(s, notificatorServer{})
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}

	// TODO: errorf
	log.Printf("server listening at %s", lis.Addr())
}
