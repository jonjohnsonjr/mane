package trace

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/jonjohnsonjr/mane"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

func Main(ctx context.Context, w io.Writer, main func(context.Context) error) error {
	return Wrap(w, main)(ctx)
}

func Wrap(w io.Writer, f mane.Func) mane.Func {
	return func(ctx context.Context) error {
		exporter, err := stdouttrace.New(stdouttrace.WithWriter(w))
		if err != nil {
			return fmt.Errorf("creating stdout exporter: %w", err)
		}
		tp := trace.NewTracerProvider(trace.WithBatcher(exporter))
		otel.SetTracerProvider(tp)

		ctx, span := otel.Tracer("mane").Start(ctx, "main")

		mainErr := f(ctx)

		span.End()

		if err := tp.Shutdown(ctx); err != nil {
			return errors.Join(mainErr, fmt.Errorf("trace provider shutdown: %v", err))
		}

		return mainErr
	}
}
