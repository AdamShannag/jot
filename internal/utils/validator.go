package utils

import (
	"regexp"
	"strings"
)

type Validator interface {
	ValidatePath(path string) bool
	ValidateName(path string) bool
}

type DefaultValidator struct{}

func (*DefaultValidator) ValidatePath(path string) bool {
	if !strings.HasSuffix(path, "/") {
		return false
	}
	pattern, _ := regexp.Compile(`^\.{0,2}/((\w/?)|(\.{2}/))*$`)
	return pattern.MatchString(path)
}

func (*DefaultValidator) ValidateName(name string) bool {
	pattern, _ := regexp.Compile(`^(\w-?)+$`)
	return pattern.MatchString(name)
}
