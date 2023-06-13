package pprof

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"runtime/pprof"
)

func Main(ctx context.Context, filename string, main func(context.Context) error) error {
	file, err := createFile(filename)
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

func createFile(filename string) (*os.File, error) {
	if filename != "" {
		return os.Create(filename)
	}

	return os.CreateTemp("", "mane-pprof")
}
