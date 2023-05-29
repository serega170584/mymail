package app

import (
	proto "awesomeProject/internal/proto"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"awesomeProject/internal/service"
)

type App struct {
	Config *viper.Viper
	Tracer trace.Tracer
}

func New(config *viper.Viper, tracer trace.Tracer) *App {
	return &App{
		config,
		tracer,
	}
}

func (app *App) Run(ctx context.Context) error {

	_, span := app.Tracer.Start(ctx, "app.run")
	defer span.End()

	const NetworkLayerTypeTcp = "tcp"

	lis, err := net.Listen(
		NetworkLayerTypeTcp,
		fmt.Sprintf("%s:%s", app.Config.GetString("app.host"), app.Config.GetString("app.port")),
	)
	if err != nil {
		log.Printf("failed to listen: %s", err.Error())
		return err
	}

	s := grpc.NewServer()
	notificatorServer := service.New(app.Config, app.Tracer)
	proto.RegisterNotificatorServer(s, notificatorServer)
	log.Printf("server listening at %s", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Printf("failed to serve: %s", err.Error())
		return err
	}

	return nil
}
