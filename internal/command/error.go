package command

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func onUsageError(cCtx *cli.Context, err error, isSubcommand bool) error {
	fmt.Fprintf(cCtx.App.Writer, "an error has occurred\n")
	return err
}

func Error(err error) error {
	return errors.New(color.RedString(err.Error()))
}
