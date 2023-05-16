package utils

import (
	"strings"
)

const defaultProjectName = "project-name"

type Formatter interface {
	FormatName(name *string)
	FormatPath(path *string)
}

type DefaultFormatter struct{}

func (*DefaultFormatter) FormatName(name *string) {
	if *name == "" {
		*name = defaultProjectName
	}
}

func (*DefaultFormatter) FormatPath(path *string) {
	*path = strings.Replace(*path, "\\", "/", -1)
	if *path == "" || *path == "." {
		*path = "./"
	} else {
		if !(strings.HasPrefix(*path, ".") ||
			strings.HasPrefix(*path, "/")) {
			*path = "/" + *path
		}
		if !strings.HasSuffix(*path, "/") {
			*path = *path + "/"
		}
	}
}
