package middleware

import (
	"fmt"
	"log"
	"testing"

	"github.com/AdamShannag/jot/internal/command/path"
	"github.com/AdamShannag/jot/internal/io"
	"github.com/AdamShannag/jot/internal/makefile"
	"github.com/AdamShannag/jot/internal/spinner"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/spf13/afero"
)

func Test_Create_DefaultMiddleware(t *testing.T) {
	srv := "user-service"

	appFs := io.SwitchToMemMap()
	services := []types.Service{
		{Name: srv},
	}

	s := types.NewSpecs("test", services, nil)
	mk := makefile.New(path.Path(path.GoModPath, srv), 10, spinner.New("cyan"))

	Defualt("Test", s, srv, mk)

	paths := []string{
		fmt.Sprintf(path.DefaultMiddlewareDirPath+"test.go", srv),
	}

	for _, path := range paths {
		fileExists(t, &appFs, path)
	}
}

func fileExists(t *testing.T, appFs *afero.Fs, path string) {
	if ok, err := afero.Exists(*appFs, path); err != nil {
		log.Panic(err)
	} else if !ok {
		t.Errorf("Expected file to exist at %s, but it is not", path)
	}
}
