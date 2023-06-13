package trace

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Main spins up otel-desktop-viewer and an otel exporter.
func Main(ctx context.Context, main func(context.Context) error) error {
	cmd := exec.CommandContext(ctx, "otel-desktop-viewer")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4318")
	os.Setenv("OTEL_TRACES_EXPORTER", "otlp")
	os.Setenv("OTEL_EXPORTER_OTLP_PROTOCOL", "http/protobuf")

	return Context(ctx, main)
}

// Context spins up an otel exporter.
func Context(ctx context.Context, main func(context.Context) error) error {
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
