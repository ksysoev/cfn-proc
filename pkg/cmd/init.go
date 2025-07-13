package cmd

import (
	"github.com/spf13/cobra"
)

type BuildInfo struct {
	Version string
}

type cmdArgs struct {
	Version string
	Out     string
}

// InitCommand initializes and returns a cobra.Command for CloudFormation pre-processing tasks.
// It configures the command with build information and sets up necessary flags.
// Accepts build that contains application version information.
// Returns a fully configured cobra.Command instance.
func InitCommand(build BuildInfo) cobra.Command {
	args := cmdArgs{
		Version: build.Version,
	}

	cmd := cobra.Command{
		Use:   "cfn-proc",
		Short: "CloudFormation pre-processor",
		Long:  "cfn-proc is a tool for processing CloudFormation templates, allowing for custom pre-processing steps before deployment.",
		RunE: func(cmd *cobra.Command, sargs []string) error {
			return runProcessor(cmd.Context(), args, sargs)
		},
	}

	cmd.Flags().StringVarP(&args.Out, "out", "o", "output.yml", "Output file")

	return cmd
}
