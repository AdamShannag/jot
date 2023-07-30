package project

import (
	"net/http"
	"testing"

	"github.com/AdamShannag/jot/v2/api/endpoint"
	"github.com/AdamShannag/jot/v2/api/endpoint/url"
	"github.com/AdamShannag/jot/v2/api/middleware"
	"github.com/AdamShannag/jot/v2/api/service"
	"github.com/AdamShannag/jot/v2/types/model"
)

func Test_ProjectBuilder(t *testing.T) {
	tests := []struct {
		name        string
		testProject model.Project
	}{
		{"Building a project", model.Project{
			Name: "Test",
			Services: []model.Service{
				{
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
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := url.NewBuilder().Path("/test").Handler("Test").Method(http.MethodGet).Build()
			endpoint := endpoint.NewBuilder().Name("Test").Urls([]model.Url{url}).Build()
			middleware := middleware.NewBuilder().Name("Test").Build()

			service := service.NewBuilder().Name("Test").
				Port(9090).
				Endpoints([]model.Endpoint{endpoint}).
				Middlewares([]model.Middleware{middleware}).
				Build()

			project := NewBuilder().Name("Test").Services([]model.Service{service}).Build()

			actualService := project.Services[0]
			expectedService := tt.testProject.Services[0]
			actualEndpoint := actualService.Endpoints[0]
			expectedEndpoint := expectedService.Endpoints[0]

			if project.Name != tt.testProject.Name {
				t.Errorf("Expected project name %q, got %q", tt.testProject.Name, project.Name)
			}

			assertEquals[string](t, tt.testProject.Name, project.Name, "project name")

			assertEquals[string](t, expectedService.Name, actualService.Name, "service name")
			assertEquals[int](t, expectedService.Port, actualService.Port, "service port")

			assertEquals[string](t, expectedService.Name, actualEndpoint.Name, "endpoint name")
			assertEquals[string](t, expectedEndpoint.Urls[0].Path, actualEndpoint.Urls[0].Path, "url path")
			assertEquals[string](t, expectedEndpoint.Urls[0].Handler, actualEndpoint.Urls[0].Handler, "url handler")
			assertEquals[string](t, expectedEndpoint.Urls[0].Method, actualEndpoint.Urls[0].Method, "url method")

			assertEquals[string](t, expectedService.Middlewares[0].Name, actualService.Middlewares[0].Name, "middleware name")

		})
	}
}

func assertEquals[V string | int](t *testing.T, expected V, actual V, property string) {
	if expected != actual {
		t.Errorf("Expected %s %q, got %q", property, expected, actual)
	}
}
