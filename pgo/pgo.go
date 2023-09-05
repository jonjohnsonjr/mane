package pgo

import (
	"context"
	"os"
	"runtime/pprof"
)

func Main(ctx context.Context, main func(context.Context) error) error {
	file, err := os.Create("default.pgo")
	if err != nil {
		return err
	}
	defer file.Close()

	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	return main(ctx)
}
