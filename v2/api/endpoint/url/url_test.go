package url

import (
	"net/http"
	"testing"

	"github.com/AdamShannag/jot/v2/types/model"
)

func Test_UrlBuilder(t *testing.T) {
	tests := []struct {
		name    string
		testUrl model.Url
	}{
		{"Building Url", model.Url{Path: "/test", Handler: "Test", Method: http.MethodGet}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := NewBuilder().Path("/test").Handler("Test").Method(http.MethodGet).Build()

			assertEquals[string](t, tt.testUrl.Path, url.Path, "url path")
			assertEquals[string](t, tt.testUrl.Handler, url.Handler, "url handler")
			assertEquals[string](t, tt.testUrl.Method, url.Method, "url method")
		})
	}
}

func assertEquals[V string | int](t *testing.T, expected V, actual V, property string) {
	if expected != actual {
		t.Errorf("Expected %s %q, got %q", property, expected, actual)
	}
}
