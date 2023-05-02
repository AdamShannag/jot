package log

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	CREATED = iota
	IGNORED
	FAILED
)

func Info(source string, action int) {
	sr := color.CyanString(source)
	switch action {
	case CREATED:
		fmt.Printf("[%s] %s\n", sr, color.GreenString("Created!"))
	case IGNORED:
		fmt.Printf("[%s] %s\n", sr, color.YellowString("Ignored!"))
	case FAILED:
		fmt.Printf("[%s] %s\n", sr, color.RedString("Failed!"))
	}
}
