package main

import (
	"awesomeProject/internal"
	pb "awesomeProject/proto"
	"context"
	"encoding/json"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

var to = flag.String("to", "test", "To name")

func main() {
	config := &internal.Config{}

	file, err := os.Open("../config-local.json")
	if err != nil {
		log.Fatalf("failed to open config: %v", err)
	}
	defer file.Close()

	d := json.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		log.Fatalf("failed to decode config: %v", err)
	}

	flag.Parse()
	conn, err := grpc.Dial(config.App.Host+":"+config.App.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
