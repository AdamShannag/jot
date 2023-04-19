package new

import (
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func Command(name string, aliases []string, category, usage, usageText, description string, flags []cli.Flag, action cli.ActionFunc, onUsageError cli.OnUsageErrorFunc) *cli.Command {
	return &cli.Command{
		Name:            name,
		Aliases:         aliases,
		Category:        category,
		Usage:           usage,
		UsageText:       usageText,
		Description:     description,
		Flags:           flags,
		SkipFlagParsing: false,
		HideHelp:        false,
		Hidden:          false,
		Action:          action,
		OnUsageError:    onUsageError,
		Before: func(cCtx *cli.Context) error {
			return nil
		},
		After: func(ctx *cli.Context) error {
			color.Cyan("Done!")
			return nil
		},
	}
}

func BoolFlag(name string, value bool, usage string, required bool, aliases ...string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:     name,
		Value:    value,
		Usage:    usage,
		Aliases:  aliases,
		Required: required,
	}
}

func IntFlag(name string, value int, usage string, required bool, aliases ...string) *cli.IntFlag {
	return &cli.IntFlag{
		Name:     name,
		Value:    value,
		Usage:    usage,
		Aliases:  aliases,
		Required: required,
	}
}

func StringFlag(name, value, usage string, required bool, aliases ...string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:     name,
		Value:    value,
		Usage:    usage,
		Aliases:  aliases,
		Required: required,
	}
}

func StringSliceFlag(name, value, usage string, required bool, aliases ...string) *cli.StringSliceFlag {
	return &cli.StringSliceFlag{
		Name:     name,
		Value:    &cli.StringSlice{},
		Usage:    usage,
		Aliases:  aliases,
		Required: required,
	}
}
