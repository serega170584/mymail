package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"

	"awesomeProject/internal/jaeger"
	"awesomeProject/internal/logger"
	notificator "awesomeProject/internal/proto"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/gomail.v2"
)

type NotificatorServer struct {
	notificator.UnimplementedNotificatorServer
	config *viper.Viper
	logger *logger.Logger
}

func New(config *viper.Viper, logger *logger.Logger) *NotificatorServer {
	return &NotificatorServer{
		config: config,
		logger: logger,
	}
}

func (server *NotificatorServer) Email(ctx context.Context, in *notificator.EmailRequest) (*emptypb.Empty, error) {
	tracerProvider, err := jaeger.TracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		server.logger.Error(fmt.Sprintf("failed to init jaeger: %s\n", err.Error()))
	}
	otel.SetTracerProvider(tracerProvider)
	tracer := tracerProvider.Tracer("notif-server-email")
	_, span := tracer.Start(ctx, "Test")
	defer span.End()

	server.logger.Info(fmt.Sprintf("To: %s, Subject: %s", in.GetTo(), in.GetSubject()))

	message := gomail.NewMessage()
	message.SetHeader("From", server.config.GetString("mail.from"))
	message.SetHeader("To", in.To...)
	message.SetHeader("Subject", fmt.Sprintf("grpc handler was triggered at %s", time.Now().String()))

	// TODO: google mailchimp если сложно то найдем другое решение
	dialer := gomail.NewDialer(server.config.GetString("mail.host"), server.config.GetInt("mail.port"), server.config.GetString("mail.from"), "111")
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := dialer.DialAndSend(message); err != nil {
		server.logger.Error(fmt.Sprintf("failed to send mail: %s\n", err.Error()))
		return &emptypb.Empty{}, err
	}
	server.logger.Info("Letter is sent")

	return &emptypb.Empty{}, nil
}

func (server *NotificatorServer) Sms(ctx context.Context, in *notificator.SmsRequest) (*emptypb.Empty, error) {
	log.Printf("method is not implemented, exit code 1")

	return &emptypb.Empty{}, nil
}
