package cmd

import (
	"context"
	"fmt"

	"github.com/ksysoev/cfn-proc/pkg/file"
)

// runProcessor processes the input file specified in sargs and writes the transformed output to the file specified in args.Out.
// It validates that exactly one input file is provided in sargs and fails if input and output paths are the same.
// Returns an error if the input file cannot be opened, decoded, or if writing the output file fails.
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
