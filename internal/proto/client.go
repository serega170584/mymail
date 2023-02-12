package notifier

import (
	"awesomeProject/internal/config"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

// TODO: golangci-lint
func main() {
	mainConfig, err := config.NewConfig()
	if err != nil {
		log.Printf("Config handle error: %v", err)
		return
	}

	// TODO: fmt.sprintf
	conn, err := grpc.Dial(mainConfig.App.Host+":"+mainConfig.App.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	c := __.NewMyMailSerivceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Email(ctx, &EmailRequest{To: []string{mainConfig.Mail.To}})
	if err != nil {
		log.Printf("could not greet: %v", err)
		return
	}
	log.Printf("Sending: %s", r.String())
}
