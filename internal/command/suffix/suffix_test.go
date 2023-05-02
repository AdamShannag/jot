package suffix

import "testing"

func Test_ServiceSuffix(t *testing.T) {
	expected := "user-service"
	actual := "user"

	ServiceSuffix(&actual)

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actual = "user-service"

	ServiceSuffix(&actual)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func Test_DockerFileSuffix(t *testing.T) {
	expected := "user-service.dockerfile"
	actual := DockerfileSuffix("user-service")

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func Test_AppSuffix(t *testing.T) {
	expected := "userApp"
	actual := AppSuffix("user")

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func Test_GoSuffix(t *testing.T) {
	expected := "userApp.go"
	actual := GoSuffix("userApp")

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
