package pprof

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"runtime/pprof"
)

func Main(ctx context.Context, main func(context.Context) error) error {
	file, err := os.CreateTemp("", "mane-pprof")
	if err != nil {
		return err
	}

	pprof.StartCPUProfile(file)
	mainErr := main(ctx)
	pprof.StopCPUProfile()

	cmd := exec.CommandContext(ctx, "go", "tool", "pprof", "-http=:", file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return errors.Join(mainErr, cmd.Run())
}
