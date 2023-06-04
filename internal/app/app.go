package app

import (
	"awesomeProject/internal/event"
	proto "awesomeProject/internal/proto"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"awesomeProject/internal/service"
)

type App struct {
	Config *viper.Viper
	Tracer trace.Tracer
	Event  *event.Event
}

func New(config *viper.Viper, tracer trace.Tracer, e *event.Event) *App {
	return &App{
		config,
		tracer,
		e,
	}
}

func (app *App) Run(ctx context.Context) error {

	_, span := app.Tracer.Start(ctx, "app.run")
	defer span.End()

	const NetworkLayerTypeTcp = "tcp"

	lis, err := net.Listen(
		NetworkLayerTypeTcp,
		net.JoinHostPort(app.Config.GetString("app.host"), app.Config.GetString("app.port")),
	)
	if err != nil {
		log.Printf("failed to listen: %s", err.Error())
		return err
	}

	s := grpc.NewServer()
	notificatorServer := service.New(app.Config, app.Tracer, app.Event)
	proto.RegisterNotificatorServer(s, notificatorServer)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		ch := <-sigCh
		fmt.Printf("got signal %v, attempting graceful shutdown", ch)
		s.GracefulStop()
	}()

	log.Printf("server listening at %s", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Printf("failed to serve: %s", err.Error())
		return err
	}

	return nil
}
