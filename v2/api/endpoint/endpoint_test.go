package endpoint

import (
	"net/http"
	"testing"

	"github.com/AdamShannag/jot/v2/api/endpoint/url"
	"github.com/AdamShannag/jot/v2/types/model"
)

func Test_EndpointBuilder(t *testing.T) {
	tests := []struct {
		name         string
		testEndpoint model.Endpoint
	}{
		{"Building an endpoint",
			model.Endpoint{
				Name: "Test",
				Urls: []model.Url{
					{
						Path:    "/test",
						Handler: "Test",
						Method:  http.MethodGet,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := url.NewBuilder().Path("/test").Handler("Test").Method(http.MethodGet).Build()
			endpoint := NewBuilder().Name("Test").Urls([]model.Url{url}).Build()

			assertEquals[string](t, tt.testEndpoint.Name, endpoint.Name, "endpoint name")
			assertEquals[string](t, tt.testEndpoint.Urls[0].Path, endpoint.Urls[0].Path, "url path")
			assertEquals[string](t, tt.testEndpoint.Urls[0].Handler, endpoint.Urls[0].Handler, "url handler")
			assertEquals[string](t, tt.testEndpoint.Urls[0].Method, endpoint.Urls[0].Method, "url method")
		})
	}
}

func assertEquals[V string | int](t *testing.T, expected V, actual V, property string) {
	if expected != actual {
		t.Errorf("Expected %s %q, got %q", property, expected, actual)
	}
}
