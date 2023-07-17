package template

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"text/template"

	"github.com/AdamShannag/jot/v1/internal/io"
	"github.com/spf13/afero"
)

func Test_Create(t *testing.T) {
	appFs := io.SwitchToMemMap()

	path := "./"
	filename := "test.test"
	pathToFile := fmt.Sprintf("%s%s", path, filename)
	tplfile := "test.test.gotpl"
	tplfilePath := "./test_data/test.test.gotpl"

	tmpl = template.Must(template.ParseFiles(tplfilePath))

	data := map[string]string{
		"Data": "some data",
	}

	Create(tplfile, path, filename, data)

	// test if file is created
	if ok, err := afero.Exists(appFs, pathToFile); err != nil {
		log.Panic(err)
	} else if !ok {
		t.Errorf("Expected file to exist at %s, but it is not", pathToFile)
	}

	// check file content
	bs, err := afero.ReadFile(appFs, pathToFile)
	if err != nil {
		log.Panic(err)
	}
	if !strings.Contains(data["Data"], string(bs)) {
		t.Errorf("Expected file to contain [%s], but got [%s]", data["Data"], string(bs))
	}

}
