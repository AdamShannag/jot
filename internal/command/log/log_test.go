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

	if !strings.Contains(output, "[Test] Created!") {
		t.Errorf("Expected to find [Test] Cerated!, but it is not there")
	}
	if !strings.Contains(output, "[Test] Ignored!") {
		t.Errorf("Expected to find [Test] Ignored, but it is not there")
	}
	if !strings.Contains(output, "[Test] Failed!") {
		t.Errorf("Expected to find [Test] Failed!, but it is not there")
	}
}
