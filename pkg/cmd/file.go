package cmd

import (
	"context"
	"fmt"

	"github.com/ksysoev/cfn-proc/pkg/file"
)

func runProcessor(_ context.Context, args cmdArgs, sargs []string) error {
	if len(sargs) != 1 {
		return fmt.Errorf("expected one input file")
	}

	processor, err := file.New(sargs[0], args.Out)
	if err != nil {
		return err
	}

	if err := processor.Process(); err != nil {
		return err
	}

	return nil
}
