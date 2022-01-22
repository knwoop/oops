package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"

	"github.com/knwoop/oops/cmd"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %s\n", err.Error())
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), unix.SIGINT, unix.SIGTERM)
	defer cancel()

	return cmd.Execute(ctx)
}
