package main

import (
	"context"
	"os"

	"github.com/AdamShannag/jot/v1/internal/cleanup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		run(os.Args)
		cancel()
	}()
	cleanup.Clean(ctx)
}
