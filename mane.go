package mane

import (
	"context"
	"log"
)

// Main is like main() but with a signal-respecting ctx and error return.
// It will log.Fatal if the given function returns an error.
func Main(main func(ctx context.Context) error) {
	if err := main(Context()); err != nil {
		log.Fatal(err)
	}
}
