package main

import (
	pb "awesomeProject/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedMyMailSerivceServer
}

func (s *server) MyMail(ctx context.Context, in *pb.MyMailRequest) (*emptypb.Empty, error) {
	log.Printf("Received: %v", in.GetTo()+" "+in.GetSubject())
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4444")
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
