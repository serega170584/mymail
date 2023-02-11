package main

import (
	"awesomeProject/internal"
	pb "awesomeProject/proto"
	"context"
	"crypto/tls"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/gomail.v2"
	"log"
	"net"
	"os"
	"time"
)

type server struct {
	pb.UnimplementedMyMailSerivceServer
}

func (s *server) MyMail(ctx context.Context, in *pb.MyMailRequest) (*emptypb.Empty, error) {
	log.Printf("Received: %v", in.GetTo()+" "+in.GetSubject())

	config := &internal.Config{}

	file, err := os.Open("../config/config-local.json")
	if err != nil {
		log.Printf("failed to open config: %v", err)
		return &emptypb.Empty{}, err
	}
	defer file.Close()

	d := json.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		log.Printf("failed to decode config: %v", err)
		return &emptypb.Empty{}, err
	}

	message := gomail.NewMessage()

	message.SetHeader("From", config.Mail.From)
	message.SetHeader("To", config.Mail.To)
	message.SetHeader("Subject", "grpc handler was triggered at"+time.Now().String())

	dialer := gomail.NewDialer(config.Mail.Host, config.Mail.Port, config.Mail.From, config.Mail.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		log.Printf("failed to send mail: %v", err)
		return &emptypb.Empty{}, err
	}
	log.Printf("Letter is sent")

	return &emptypb.Empty{}, nil
}

func main() {
	config := &internal.Config{}

	file, err := os.Open("../config/config-local.json")
	if err != nil {
		log.Fatalf("failed to open config: %v", err)
	}
	defer file.Close()

	d := json.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		log.Printf("failed to decode config: %v", err)
		return
	}

	lis, err := net.Listen("tcp", ":"+config.App.Port)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterMyMailSerivceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
