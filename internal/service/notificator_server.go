package notificator

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	notificator "awesomeProject/internal/proto"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/gomail.v2"
)

type NotificatorServer struct {
	notificator.UnimplementedNotificatorServer
	config *viper.Viper
}

func New(config *viper.Viper) *NotificatorServer {
	return &NotificatorServer{
		config: config,
	}
}

func (server *NotificatorServer) Email(ctx context.Context, in *notificator.EmailRequest) (*emptypb.Empty, error) {
	log.Printf("To: %s, Subject: %s\n", in.GetTo(), in.GetSubject())

	message := gomail.NewMessage()
	message.SetHeader("From", server.config.GetString("mail.from"))
	message.SetHeader("To", in.To...)
	message.SetHeader("Subject", fmt.Sprintf("grpc handler was triggered at %s", time.Now().String()))

	// TODO: google mailchimp если сложно то найдем другое решение
	dialer := gomail.NewDialer(server.config.GetString("mail.host"), server.config.GetInt("mail.port"), server.config.GetString("mail.from"), "111")
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := dialer.DialAndSend(message); err != nil {
		log.Printf("failed to send mail: %s\n", err.Error())
		return &emptypb.Empty{}, err
	}
	log.Println("email is sent")

	return &emptypb.Empty{}, nil
}

func (server *NotificatorServer) Sms(ctx context.Context, in *notificator.SmsRequest) (*emptypb.Empty, error) {
	log.Println("method is not implemented, exit code 1")

	return &emptypb.Empty{}, nil
}
