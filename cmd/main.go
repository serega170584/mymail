package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/proto"
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/gomail.v2"
	"log"
	"net"
	"time"
)

type server struct {
	__.UnimplementedMyMailSerivceServer
}

func (s *server) MyMail(ctx context.Context, in *__.MyMailRequest) (*emptypb.Empty, error) {
	log.Printf("Received: %v", in.GetTo()+" "+in.GetSubject())

	mainConfig, err := config.NewConfig()
	if err != nil {
		log.Printf("Config handle error: %v", err)
		return &emptypb.Empty{}, err
	}

	message := gomail.NewMessage()

	message.SetHeader("From", mainConfig.Mail.From)
	message.SetHeader("To", mainConfig.Mail.To)
	message.SetHeader("Subject", "grpc handler was triggered at"+time.Now().String())

	dialer := gomail.NewDialer(mainConfig.Mail.Host, mainConfig.Mail.Port, mainConfig.Mail.From, mainConfig.Mail.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		log.Printf("failed to send mail: %v", err)
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

	lis, err := net.Listen("tcp", "localhost"+":"+mainConfig.App.Port)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	__.RegisterMyMailSerivceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
}
