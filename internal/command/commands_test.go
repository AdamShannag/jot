package command

import "testing"

func Test_addCommands(t *testing.T) {
	addCommands(ini())

	if len(commands) != 1 {
		t.Errorf("Expected to find one command, but found %d", len(commands))
	}
}
