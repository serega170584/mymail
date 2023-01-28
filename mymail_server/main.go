package main

import (
	pb "awesomeProject/proto"
	"context"
	"log"
)

type server struct {
	pb.UnimplementedMyMailSerivceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.MyMailRequest) (*pb.My, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {

}
