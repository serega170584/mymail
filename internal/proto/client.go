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
	mainConfig := config.NewConfig()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", mainConfig.GetString("app.host"), mainConfig.GetString("app.port")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %s", err.Error())
		return
	}
	defer conn.Close()
	c := NewNotificatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Email(ctx, &EmailRequest{To: mainConfig.GetStringSlice("mail.to")})
	if err != nil {
		log.Printf("could not greet: %s", err.Error())
		return
	}
	log.Printf("Sending: %s", r.String())
}
