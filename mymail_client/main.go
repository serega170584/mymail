package main

import (
	pb "awesomeProject/proto"
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/gomail.v2"
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

	message := gomail.NewMessage()

	message.SetHeader("From", config.App.From)
	message.SetHeader("To", config.App.To)
	message.SetHeader("Subject", "grpc handler was triggered at"+time.Now().String())

	dialer := gomail.NewDialer(config.Mail.Host, config.Mail.Port, "from@gmail.com", "<email_password>")
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		log.Fatalf("failed to send mail: %v", err)
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
