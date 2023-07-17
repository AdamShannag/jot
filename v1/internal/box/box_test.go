package box

import (
	"strings"
	"testing"
)

func Test_Welcome(t *testing.T) {
	msg := Welcome()

	if !strings.Contains(msg, "Jot") {
		t.Error("Expected to find 'Jot', but it is not there")
	}
}
