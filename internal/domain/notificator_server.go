package domain

import (
	"awesomeProject/internal/config"
	notificator "awesomeProject/internal/proto"
	"crypto/tls"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/gomail.v2"

	"context"
	"fmt"
	"log"
	"time"
)

type NotificatorServer struct {
	notificator.UnimplementedNotificatorServer
}

func NewNotificatorServer() *NotificatorServer {
	return &NotificatorServer{}
}

func (r *NotificatorServer) Email(ctx context.Context, in *notificator.EmailRequest) (*emptypb.Empty, error) {
	mainConfig := config.NewConfig()

	log.Println("Received:", in.GetTo(), in.GetSubject())

	message := gomail.NewMessage()

	message.SetHeader("From", mainConfig.GetString("mail.from"))
	message.SetHeader("To", in.To...)
	message.SetHeader("Subject", fmt.Sprintf("grpc handler was triggered at %s", time.Now().String()))

	// TODO: google mailchimp если сложно то найдем другое решение
	dialer := gomail.NewDialer(mainConfig.GetString("mail.host"), mainConfig.GetInt("mail.port"), mainConfig.GetString("mail.from"), "111")
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		// todo назвать различия прописывания ошибки двумя способами
		log.Printf("failed to send mail: %v\n", err)
		log.Printf("failed to send mail: %s\n", err.Error())
		return &emptypb.Empty{}, err
	}
	log.Printf("Letter is sent")

	return &emptypb.Empty{}, nil
}

func (r *NotificatorServer) Sms(ctx context.Context, in *notificator.SmsRequest) (*emptypb.Empty, error) {
	log.Printf("Letter is sent")

	return &emptypb.Empty{}, nil
}
