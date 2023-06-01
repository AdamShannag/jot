package log

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Info(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	Info("Test", CREATED)
	Info("Test", IGNORED)
	Info("Test", FAILED)

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "CREATE Test") {
		t.Errorf("Expected to find Create Test, but it is not there")
	}
	if !strings.Contains(output, "IGNORE Test") {
		t.Errorf("Expected to find IGNORE Test, but it is not there")
	}
	if !strings.Contains(output, "FAIL Test") {
		t.Errorf("Expected to find FAIL Test, but it is not there")
	}
}
