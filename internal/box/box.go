package box

import (
	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/fatih/color"
)

const (
	header = "Generate microservices-architecture projects"
	github = "https://github.com/AdamShannag/jot"
)

func Welcome() {
	Box := box.New(
		box.Config{
			Px:           5,
			Py:           4,
			Type:         "Double",
			Color:        "Cyan",
			ContentAlign: "Center",
		},
	)
	Box.Print("Jot", fmt.Sprintf("%s\n\n%s", header, color.BlueString(github)))
}
