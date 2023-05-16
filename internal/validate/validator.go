package validate

import (
	"errors"
	"regexp"
	"strings"
)

func Path(path string) error {
	pattern, _ := regexp.Compile(`^\.{0,2}/((\w/?)|(\.{2}/))*$`)
	if !strings.HasSuffix(path, "/") || !pattern.MatchString(path) {
		return errors.New("invalid path project")
	}
	return nil
}

func Name(name string) error {
	pattern, _ := regexp.Compile(`^(\w-?)+$`)
	if !pattern.MatchString(name) {
		return errors.New("invalid name project")
	}
	return nil
}
