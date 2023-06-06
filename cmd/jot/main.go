package main

import (
	"context"
	"os"

	"github.com/AdamShannag/jot/internal/cleanup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		run(os.Args)
		cancel()
	}()
	cleanup.Clean(ctx)
}
