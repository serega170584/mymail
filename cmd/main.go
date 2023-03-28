package main

import (
	"context"
	"log"

	"awesomeProject/internal/app"
	"awesomeProject/internal/config"
	"awesomeProject/internal/jaeger"
)

func main() {
	ctx := context.Background()
	tracer := jaeger.New(ctx)

	ctx, span := tracer.Start(ctx, "mymail-main")
	defer span.End()

	appApp := app.New(config.New(), tracer)
	err := appApp.Run(ctx)
	if err != nil {
		log.Printf("server is interrupted, err: %s\n", err.Error())
	}
}
