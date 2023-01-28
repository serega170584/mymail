package main

import (
	pb "awesomeProject/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 4444, "The server port")
)

// server is used to implement GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.MyMailRequest) ([]string, error) {
	log.Printf("Received: %v", in.GetTo())
	return in.GetTo(), nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
