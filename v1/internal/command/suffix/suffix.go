package suffix

import (
	"fmt"
	"strings"

	p "github.com/AdamShannag/jot/v1/internal/command/path"
)

func ServiceSuffix(name string) string {
	if !strings.HasSuffix(name, "-service") {
		return name + "-service"
	}
	return name
}

func DockerfileSuffix(filename string) string {
	return fmt.Sprintf("%s%s", filename, p.DockerImageFileExt)

}

func AppSuffix(s string) string {
	return fmt.Sprintf("%sApp", s)

}

func GoSuffix(s string) string {
	return fmt.Sprintf("%s.go", s)
}
