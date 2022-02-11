package solution

import (
	"strings"
	"testing"
)

func TestGetMessage(t *testing.T) {
	actual := GetMessage()
	expect := "Hello 🗺️!"

	if !strings.EqualFold(actual, expect) {
		t.Errorf("actual %q expect %q", actual, expect)
	}

	t.Log("OK!")
}
