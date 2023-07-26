package writer

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/AdamShannag/jot/v2/api/endpoint"
	"github.com/AdamShannag/jot/v2/api/endpoint/url"
	"github.com/AdamShannag/jot/v2/api/middleware"
	"github.com/AdamShannag/jot/v2/api/project"
	"github.com/AdamShannag/jot/v2/api/service"
	"github.com/AdamShannag/jot/v2/types/model"
	"github.com/AdamShannag/jot/v2/writer/fs"
	"github.com/spf13/afero"
)

const testDir = "./test-data/"

func TestProjectWriter(t *testing.T) {
	url := url.NewBuilder().Path("/test").Handler("Test").Method(http.MethodGet).Build()
	endpoint := endpoint.NewBuilder().Name("test").Urls([]model.Url{url}).Build()
	middleware := middleware.NewBuilder().Defualt("test").Build()

	service := service.NewBuilder().Name("service").
		Port(9090).
		Endpoints([]model.Endpoint{endpoint}).
		Middlewares([]model.Middleware{middleware}).
		Build()

	project := project.NewBuilder().Name("my-project").Services([]model.Service{service}).Build()

	os.Mkdir(testDir, os.ModePerm)

	w := NewProjectWriter(project)
	w.Write(testDir)

	assertExists(t, "my-project")
	assertExists(t, "my-project/service")
	assertExists(t, "my-project/service/api")
	assertExists(t, "my-project/service/cmd")
	assertExists(t, "my-project/service/bin")
	assertExists(t, "my-project/service/deploy")
	assertExists(t, "my-project/service/api/endpoints")
	assertExists(t, "my-project/service/api/endpoints/test")
	assertExists(t, "my-project/service/api/endpoints/test/test.go")
	assertExists(t, "my-project/service/api/middlewares")
	assertExists(t, "my-project/service/api/middlewares/test")
	assertExists(t, "my-project/service/api/middlewares/test/test.go")
	assertExists(t, "my-project/service/cmd/service")
	assertExists(t, "my-project/service/cmd/service/main.go")

	os.RemoveAll(testDir)

}

func assertExists(t *testing.T, path string) {
	if !existsAt(testDir + path) {
		t.Errorf("Expecting directory or file [%s] to exist but it does not", path)
	}
}

func existsAt(path string) bool {
	ok, err := afero.Exists(fs.Get(), path)
	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}
	return ok
}
