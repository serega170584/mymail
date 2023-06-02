package service

import (
	"awesomeProject/internal/event"
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"log"

	notificator "awesomeProject/internal/proto"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NotificatorServer struct {
	notificator.UnimplementedNotificatorServer
	config *viper.Viper
	tracer trace.Tracer
	event  *event.Event
}

func New(config *viper.Viper, tracer trace.Tracer, e *event.Event) *NotificatorServer {
	return &NotificatorServer{
		config: config,
		tracer: tracer,
		event:  e,
	}
}

func (server *NotificatorServer) Email(ctx context.Context, in *notificator.EmailRequest) (*emptypb.Empty, error) {
	_, span := server.tracer.Start(ctx, "service-email")
	defer span.End()

	log.Printf("To: %s, Subject: %s\n", in.GetTo(), in.GetSubject())

	config := server.config

	request, err := json.Marshal(in)

	if err != nil {
		fmt.Printf("To header marshal error: %s", err.Error())
		return nil, err
	}

	err = server.event.Produce(config.GetString("events.mail_send.topic"), string(request))
	if err != nil {
		fmt.Print("Event produce error: %s", err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (server *NotificatorServer) Sms(ctx context.Context, in *notificator.SmsRequest) (*emptypb.Empty, error) {
	_, span := server.tracer.Start(ctx, "service-sms")
	defer span.End()

	log.Println("method is not implemented, exit code 1")

	return &emptypb.Empty{}, nil
}
