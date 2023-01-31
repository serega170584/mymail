package main

import (
	pb "awesomeProject/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"os"

	"gopkg.in/yaml.v3"
)

type server struct {
	pb.UnimplementedMyMailSerivceServer
}

type Config struct {
	Port string `yaml:"port"`
}

func (s *server) MyMail(ctx context.Context, in *pb.MyMailRequest) (*emptypb.Empty, error) {
	log.Printf("Received: %v", in.GetTo()+" "+in.GetSubject())
	return &emptypb.Empty{}, nil
}

func main() {
	config := &Config{}

	file, err := os.Open("../config.yaml")
	if err != nil {
		log.Fatalf("failed to open config: %v", err)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		log.Fatalf("failed to decode config: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+config.Port)
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
