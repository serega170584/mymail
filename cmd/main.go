package main

import (
	"awesomeProject/internal/event"
	"context"
	"log"
	"os"

	"awesomeProject/internal/app"
	"awesomeProject/internal/config"
	"awesomeProject/internal/jaeger"
)

const DefaultJaegerSpanName = "mymail-main"

func main() {
	ctx := context.Background()
	tracer := jaeger.New(ctx)

	spanName, ok := os.LookupEnv("JAEGER_SPAN_NAME")
	if !ok {
		spanName = DefaultJaegerSpanName
	}

	ctx, span := tracer.Start(ctx, spanName)
	defer span.End()

	e := event.New()

	appApp := app.New(config.New(), tracer, e)
	err := appApp.Run(ctx)
	if err != nil {
		log.Printf("server is interrupted, err: %s\n", err.Error())
	}
}
