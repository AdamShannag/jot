package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func onUsageError(cCtx *cli.Context, err error, isSubcommand bool) error {
	fmt.Fprintf(cCtx.App.Writer, "an error has occurred\n")
	return err
}
