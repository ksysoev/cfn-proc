package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type BuildInfo struct {
	Version string
}

type cmdArgs struct {
	Version string
}

func InitCommand(build BuildInfo) cobra.Command {
	_ = &cmdArgs{
		Version: build.Version,
	}

	cmd := cobra.Command{
		Use:   "cnf-prov",
		Short: "",
		Long:  "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return fmt.Errorf("not implemented")
		},
	}

	return cmd
}
