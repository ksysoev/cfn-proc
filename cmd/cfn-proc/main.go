package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ksysoev/cfn-proc/pkg/cmd"
)

var version = "dev"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	command := cmd.InitCommand(cmd.BuildInfo{
		Version: version,
	})

	err := command.ExecuteContext(ctx)
	cancel()
	if err != nil {
		os.Exit(1)
	}
}
