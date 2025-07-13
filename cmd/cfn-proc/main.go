package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ksysoev/cfn-proc/pkg/cmd"
)

var defaultServer = "make-it-public.dev:8081"
var version = "dev"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	command := cmd.InitCommand(cmd.BuildInfo{
		DefaultServer: defaultServer,
		Version:       version,
	})

	err := command.ExecuteContext(ctx)
	cancel()
	if err != nil {
		os.Exit(1)
	}
}
