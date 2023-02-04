package main

import (
	pb "awesomeProject/proto"
	"context"
	"crypto/tls"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"os"
	"time"
)

type server struct {
	pb.UnimplementedMyMailSerivceServer
}

type Config struct {
	App struct {
		Port string `json:"port"`
	}
	Mail struct {
		Port     int    `json:"port"`
		From     string `json:"from"`
		To       string `json:"to"`
		Host     string `json:"host"`
		Password string `json:"password"`
	}
}

func (s *server) MyMail(ctx context.Context, in *pb.MyMailRequest) (*emptypb.Empty, error) {
	log.Printf("Received: %v", in.GetTo()+" "+in.GetSubject())

	config := &Config{}

	file, err := os.Open("../config-local.json")
	if err != nil {
		log.Fatalf("failed to open config: %v", err)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		log.Fatalf("failed to decode config: %v", err)
	}

	message := gomail.NewMessage()

	message.SetHeader("From", config.Mail.From)
	message.SetHeader("To", config.Mail.To)
	message.SetHeader("Subject", "grpc handler was triggered at"+time.Now().String())

	dialer := gomail.NewDialer(config.Mail.Host, config.Mail.Port, config.Mail.From, config.Mail.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		log.Fatalf("failed to send mail: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func main() {
	config := &Config{}

	file, err := os.Open("../config-local.json")
	if err != nil {
		log.Fatalf("failed to open config: %v", err)
	}
	defer file.Close()

	d := json.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		log.Fatalf("failed to decode config: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+config.App.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyMailSerivceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
