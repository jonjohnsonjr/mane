package mane

import (
	"context"
	"os"
	"os/signal"
)

// Context handles the interrupt signal, use this instead of context.Background()
// in your main function.
func Context() context.Context {
	// TODO: Probably have a way to ask for the cancel func.
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	return ctx
}
