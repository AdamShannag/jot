package suffix

import "testing"

func Test_Contains(t *testing.T) {
	items := []string{"foo", "bar"}
	item := "foo"

	actual := Contains(items, item)

	if actual != true {
		t.Error("Expected to find item, but was not found")
	}
}

func Test_TitleCase(t *testing.T) {
	expected := "Test"
	actual := TitleCase("tEsT")

	if actual != expected {
		t.Errorf("Expected to be %s, but got %s", expected, actual)
	}
}
