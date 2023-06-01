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
		fmt.Printf("%s %s\n", color.GreenString("CREATE"), sr)
	case IGNORED:
		fmt.Printf("%s %s\n", color.YellowString("IGNORE"), sr)
	case FAILED:
		fmt.Printf("%s %s\n", color.RedString("FAIL"), sr)
	}
}
