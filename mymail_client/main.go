package main

import (
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

type Config struct {
	App struct {
		Port string `json:"port"`
		From string `json:"from"`
		To   string `json:"to"`
	}
	Mail struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}
}

func main() {
	config := &Config{}

	file, err := os.Open("../config.json")
	if err != nil {
		log.Fatalf("failed to open config: %v", err)
	}
	defer file.Close()

	d := json.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		log.Fatalf("failed to decode config: %v", err)
	}

	flag.Parse()
	conn, err := grpc.Dial("localhost:"+config.App.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
