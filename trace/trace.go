package trace

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

func Main(ctx context.Context, main func(context.Context) error) error {
	client := otlptracehttp.NewClient()
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return fmt.Errorf("creating OTLP trace exporter: %v", err)
	}
	tp := trace.NewTracerProvider(trace.WithBatcher(exporter))
	otel.SetTracerProvider(tp)

	ctx, span := otel.Tracer("mane").Start(ctx, "main")

	mainErr := main(ctx)

	span.End()

	if err := tp.Shutdown(ctx); err != nil {
		return errors.Join(mainErr, fmt.Errorf("trace provider shutdown: %v", err))
	}

	return mainErr
}
