package notifier

import (
	"awesomeProject/internal/config"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

// TODO: golangci-lint
func main() {
	mainConfig, err := config.NewConfig()
	if err != nil {
		log.Printf("Config handle error: %s", err.Error())
		return
	}

	// TODO: fmt.sprintf
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", mainConfig.App.Host, mainConfig.App.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %s", err.Error())
		return
	}
	defer conn.Close()
	c := NewNotificatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Email(ctx, &EmailRequest{To: []string{mainConfig.Mail.To}})
	if err != nil {
		log.Printf("could not greet: %s", err.Error())
		return
	}
	log.Printf("Sending: %s", r.String())
}
