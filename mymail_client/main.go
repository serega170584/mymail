package main

import (
	pb "awesomeProject/proto"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var to = flag.String("to", "test", "To name")

func main() {
	flag.Parse()
	conn, err := grpc.Dial("localhost:4444", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMyMailSerivceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.MyMail(ctx, &pb.MyMailRequest{To: *to})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Sending: %s", r.String())
}
