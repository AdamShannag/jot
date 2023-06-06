package spinner

import (
	"fmt"
	"time"

	s "github.com/briandowns/spinner"
)

type CustomSpinner struct {
	spinner *s.Spinner
}

func New(color string) *CustomSpinner {
	spinner := s.New(s.CharSets[11], 100*time.Millisecond)
	spinner.Color(color)
	spinner.Suffix = " Running..."
	return &CustomSpinner{
		spinner,
	}
}

func (c *CustomSpinner) Suffix(s string) *CustomSpinner {
	c.spinner.Restart()
	c.spinner.Suffix = fmt.Sprintf(" %s", s)
	return c
}

func (c *CustomSpinner) Start() *CustomSpinner {
	c.spinner.Start()
	return c
}

func (c *CustomSpinner) Stop() *CustomSpinner {
	c.spinner.Stop()
	return c
}
