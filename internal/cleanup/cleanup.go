package cleanup

import (
	"context"
)

var cleanupFunc []func() error

func Add(cFunc func() error) {
	cleanupFunc = append(cleanupFunc, cFunc)
}

func Clean(ctx context.Context) {
	for {
		<-ctx.Done()
		cleanAll()
		return
	}
}

func cleanAll() {
	for _, cleanup := range cleanupFunc {
		cleanup()
	}
}
