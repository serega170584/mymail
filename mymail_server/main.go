package main

import (
	pb "awesomeProject/proto"
	"context"
	"log"
)

type server struct {
	pb.UnimplementedMyMailSerivceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.MyMailRequest) (*pb.MyMailReply, error) {
	log.Printf("Received: %v", in.GetTo())
	return &pb.MyMailReply{To: "Hello " + in.GetTo()}, nil
}

func main() {

}
