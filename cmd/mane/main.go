package main

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"

	"github.com/jonjohnsonjr/mane/trace"
)

func main() {
	trace.Main(context.Background(), func(ctx context.Context) error {
		if err := foo(ctx); err != nil {
			return err
		}

		return bar(ctx)
	})
}

func foo(ctx context.Context) error {
	ctx, span := otel.Tracer("example.com/something").Start(ctx, "foo")
	defer span.End()

	time.Sleep(10 * time.Millisecond)

	return baz(ctx)
}

func bar(ctx context.Context) error {
	ctx, span := otel.Tracer("example.com/something").Start(ctx, "bar")
	defer span.End()

	time.Sleep(10 * time.Millisecond)

	return nil
}

func baz(ctx context.Context) error {
	ctx, span := otel.Tracer("example.com/something").Start(ctx, "baz")
	defer span.End()

	time.Sleep(10 * time.Millisecond)

	return nil
}