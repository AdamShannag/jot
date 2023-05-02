package suffix

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Contains(slice []string, item string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func TitleCase(s string) string {
	return cases.Title(language.English, cases.NoLower).String(strings.ToLower(s))
}
