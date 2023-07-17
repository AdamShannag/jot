package path

import "testing"

func Test_Path(t *testing.T) {
	format := "./test/%s"
	expected := "./test/app"

	actual := Path(format, "app")

	if actual != expected {
		t.Errorf("Expected to be %s, got %s", expected, actual)
	}
}
