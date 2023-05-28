package jaeger

import (
	"context"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	service           = "mymail-demo"
	environment       = "prod"
	DefaultTracerUrl  = "http://127.0.0.1:14268/api/traces"
	DefaultTracerName = "mymail-jaeger-tracer"
)

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}

	serviceName, ok := os.LookupEnv("JAEGER_APP_NAME")
	if !ok {
		serviceName = service
	}

	env, ok := os.LookupEnv("JAEGER_ENV_NAME")
	if !ok {
		env = environment
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			attribute.String("environment", env),
			attribute.String("timestamp", time.Now().String()),
		)),
	)

	return tp, nil
}

func New(ctx context.Context) trace.Tracer {
	tracerUrl, ok := os.LookupEnv("JAEGER_TRACER_URL")
	if !ok {
		tracerUrl = DefaultTracerUrl
	}

	tp, err := tracerProvider(tracerUrl)
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)

	_, cancel := context.WithCancel(ctx)
	defer cancel()

	tracerName, ok := os.LookupEnv("JAEGER_TRACER_NAME")
	if !ok {
		tracerName = DefaultTracerName
	}

	return tp.Tracer(tracerName)
}
