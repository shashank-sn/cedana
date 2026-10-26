// Cedana is automated GPU checkpointing and migration infrastructure that increases the useful work your GPUs deliver.
//
// This module holds the open source Cedana daemon and CLI, built on CRIU, which checkpoint a running Linux process or container and restore it on the same or another machine.
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/cedana/cedana/cmd"
	"github.com/cedana/cedana/pkg/version"
)

// loaded from ldflag definitions
var Version = "dev"

func main() {
	// Grandparent context to deal with OS interrupts
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	signal.Ignore(syscall.SIGPIPE) // Avoid program termination due to broken pipes

	defer stop()
	version.PutVersion(Version)

	if err := run(ctx, Version); err != nil {
		os.Exit(1)
	}
}

func run(ctx context.Context, version string) error {
	return cmd.Execute(ctx, version)
}
