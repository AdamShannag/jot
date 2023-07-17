package service

import (
	"net/http"
	"testing"

	"github.com/AdamShannag/jot/v2/api/endpoint"
	"github.com/AdamShannag/jot/v2/api/endpoint/url"
	"github.com/AdamShannag/jot/v2/api/middleware"
	"github.com/AdamShannag/jot/v2/types/model"
)

func Test_ServiceBuilder(t *testing.T) {
	tests := []struct {
		name        string
		testService model.Service
	}{
		{"Building a service", model.Service{
			Name: "Test",
			Port: 9090,
			Endpoints: []model.Endpoint{
				{
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
			Middlewares: []model.Middleware{{Name: "Test"}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := url.NewBuilder().Path("/test").Handler("Test").Method(http.MethodGet).Build()
			endpoint := endpoint.NewBuilder().Name("Test").Urls([]model.Url{url}).Build()
			middleware := middleware.NewBuilder().Defualt("Test").Build()

			service := NewBuilder().Name("Test").
				Port(9090).
				Endpoints([]model.Endpoint{endpoint}).
				Middlewares([]model.Middleware{middleware}).
				Build()

			actualEndpoint := service.Endpoints[0]
			expectedEndpoint := tt.testService.Endpoints[0]

			assertEquals[string](t, tt.testService.Name, service.Name, "service name")
			assertEquals[int](t, tt.testService.Port, service.Port, "service port")
			assertEquals[string](t, tt.testService.Endpoints[0].Name, actualEndpoint.Name, "endpoint name")
			assertEquals[string](t, expectedEndpoint.Urls[0].Path, actualEndpoint.Urls[0].Path, "url path")
			assertEquals[string](t, expectedEndpoint.Urls[0].Handler, actualEndpoint.Urls[0].Handler, "url handler")
			assertEquals[string](t, expectedEndpoint.Urls[0].Method, actualEndpoint.Urls[0].Method, "url method")
		})
	}
}

func assertEquals[V string | int](t *testing.T, expected V, actual V, property string) {
	if expected != actual {
		t.Errorf("Expected %s %q, got %q", property, expected, actual)
	}
}
