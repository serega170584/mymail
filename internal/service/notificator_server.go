package service

import (
	"awesomeProject/internal/event"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"

	"go.opentelemetry.io/otel/trace"

	notificator "awesomeProject/internal/proto"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/gomail.v2"
)

type NotificatorServer struct {
	notificator.UnimplementedNotificatorServer
	config *viper.Viper
	tracer trace.Tracer
}

func New(config *viper.Viper, tracer trace.Tracer) *NotificatorServer {
	return &NotificatorServer{
		config: config,
		tracer: tracer,
	}
}

func (server *NotificatorServer) Email(ctx context.Context, in *notificator.EmailRequest) (*emptypb.Empty, error) {
	_, span := server.tracer.Start(ctx, "service-email")
	defer span.End()

	log.Printf("To: %s, Subject: %s\n", in.GetTo(), in.GetSubject())

	eventName := "mail_send"
	eventConn, err := event.New(server.config, eventName)

	if err != nil {
		fmt.Printf("Event %s conn error: %s", eventName, err.Error())
	}
	err = eventConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Printf("EVent write deadline error: %s", err.Error())
	}
	_, err = eventConn.WriteMessages(kafka.Message{Value: []byte("test")})
	if err != nil {
		fmt.Printf("Event write message error: %s", err.Error())
	}

	message := gomail.NewMessage()
	message.SetHeader("From", server.config.GetString("mail.from"))
	message.SetHeader("To", in.To...)
	message.SetHeader("Subject", fmt.Sprintf("grpc handler was triggered at %s", time.Now().String()))

	// TODO: google mailchimp если сложно то найдем другое решение
	// dialer := gomail.NewDialer(server.config.GetString("mail.host"), server.config.GetInt("mail.port"), server.config.GetString("mail.from"), "111")
	// dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// if err := dialer.DialAndSend(message); err != nil {
	// 	log.Printf("failed to send mail: %s\n", err.Error())
	// 	return &emptypb.Empty{}, err
	// }
	log.Println("email is sent")

	return &emptypb.Empty{}, nil
}

func (server *NotificatorServer) Sms(ctx context.Context, in *notificator.SmsRequest) (*emptypb.Empty, error) {
	_, span := server.tracer.Start(ctx, "service-sms")
	defer span.End()

	log.Println("method is not implemented, exit code 1")

	return &emptypb.Empty{}, nil
}
