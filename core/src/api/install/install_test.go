package install

import (
	"testing"
)

func TestInstall(t *testing.T) {
	actual := Install("hello", "1.0")
	expected := "hello1.0"
	if expected != actual {
		t.Errorf("Error occured while testing sayhello: '%s' != '%s'", expected, actual)
	}
}
