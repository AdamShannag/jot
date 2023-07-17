package middleware

import (
	"testing"

	"github.com/AdamShannag/jot/v2/types/model"
)

func Test_MiddlewareBuilder(t *testing.T) {
	tests := []struct {
		name           string
		testMiddleware model.Middleware
	}{
		{"Building a middleware", model.Middleware{Name: "Test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			middleware := NewBuilder().Defualt("Test").Build()

			assertEquals[string](t, tt.testMiddleware.Name, middleware.Name, "middleware name")
		})
	}
}

func assertEquals[V string | int](t *testing.T, expected V, actual V, property string) {
	if expected != actual {
		t.Errorf("Expected %s %q, got %q", property, expected, actual)
	}
}
