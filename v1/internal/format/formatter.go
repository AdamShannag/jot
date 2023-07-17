package format

import (
	"strings"

	"github.com/AdamShannag/jot/v1/internal/config"
)

func ProjectName(name string) string {
	if name == "" {
		return config.DefaultProjectName
	}
	return name
}

func Path(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	if path == "" || path == "." {
		path = "./"
	} else {
		if !(strings.HasPrefix(path, ".") ||
			strings.HasPrefix(path, "/")) {
			path = "/" + path
		}
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
	}
	return path
}
