package server

import (
	"awesomeProject/internal/jaeger"
	"awesomeProject/internal/logger"
	notificator "awesomeProject/internal/proto"
	"crypto/tls"
	"go.opentelemetry.io/otel"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/gomail.v2"

	"context"
	"fmt"
	"log"
	"time"
)

type NotificatorServer struct {
	notificator.UnimplementedNotificatorServer
	config *viper.Viper
	logger *logger.Logger
}

func New(config *viper.Viper, logger *logger.Logger) *NotificatorServer {
	return &NotificatorServer{config: config, logger: logger}
}

func (server *NotificatorServer) Email(ctx context.Context, in *notificator.EmailRequest) (*emptypb.Empty, error) {
	appLogger := server.logger

	appLogger.Info(fmt.Sprintf("To: %s, Subject: %s", in.GetTo(), in.GetSubject()))

	message := gomail.NewMessage()

	config := server.config

	message.SetHeader("From", config.GetString("mail.from"))
	message.SetHeader("To", in.To...)
	message.SetHeader("Subject", fmt.Sprintf("grpc handler was triggered at %s", time.Now().String()))

	// TODO: google mailchimp если сложно то найдем другое решение
	dialer := gomail.NewDialer(config.GetString("mail.host"), config.GetInt("mail.port"), config.GetString("mail.from"), "111")
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	tracerProvider, err := jaeger.TracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		appLogger.Error(fmt.Sprintf("failed to connect jaeger: %s\n", err.Error()))
	}
	otel.SetTracerProvider(tracerProvider)
	tracer := tracerProvider.Tracer("component-main")

	ctx, span := tracer.Start(ctx, "Test")

	if err := dialer.DialAndSend(message); err != nil {
		appLogger.Error(fmt.Sprintf("failed to send mail: %s\n", err.Error()))
		return &emptypb.Empty{}, err
	}

	defer span.End()

	appLogger.Info("Letter is sent")

	return &emptypb.Empty{}, nil
}

func (r *NotificatorServer) Sms(ctx context.Context, in *notificator.SmsRequest) (*emptypb.Empty, error) {
	log.Printf("Letter is sent")

	return &emptypb.Empty{}, nil
}
