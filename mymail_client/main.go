package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/proto"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var to = flag.String("to", "test", "To name")

func main() {
	mainConfig, err := config.NewConfig()
	if err != nil {
		log.Printf("Config handle error: %v", err)
		return
	}

	flag.Parse()
	conn, err := grpc.Dial(mainConfig.App.Host+":"+mainConfig.App.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	c := __.NewMyMailSerivceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.MyMail(ctx, &__.MyMailRequest{To: *to})
	if err != nil {
		log.Printf("could not greet: %v", err)
		return
	}
	log.Printf("Sending: %s", r.String())
}
