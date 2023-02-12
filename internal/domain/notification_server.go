package domain

import (
	"awesomeProject/internal/config"
	notificator "awesomeProject/internal/proto"
	"context"
	"crypto/tls"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/gomail.v2"
	"log"
	"time"
)

type NotificatorServer struct {
	cfg config.Config

	notificator.UnimplementedNotificatorServer
}

func (r *NotificatorServer) Email(ctx context.Context, in *notificator.EmailRequest) (*emptypb.Empty, error) {
	mainConfig, err := config.NewConfig()
	if err != nil {
		log.Printf("Config handle error: %v", err)
		return &emptypb.Empty{}, err
	}

	log.Println("Received:", in.GetTo(), in.GetSubject())

	message := gomail.NewMessage()

	// tood: default in config
	message.SetHeader("From", mainConfig.Mail.From)
	message.SetHeader("To", in.To...)
	message.SetHeader("Subject", fmt.Sprintf("grpc handler was triggered at %s", time.Now().String()))

	// TODO: google mailchimp если сложно то найдем другое решение
	dialer := gomail.NewDialer(mainConfig.Mail.Host, mainConfig.Mail.Port, mainConfig.Mail.From, mainConfig.Mail.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		// todo посмотреть различия
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
