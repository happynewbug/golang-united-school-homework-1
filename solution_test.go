package solution

import (
	"strings"
	"testing"
)

func TestGetMessage(t *testing.T) {
	expected := string([]byte{72, 101, 108, 108, 111, 32, 240, 159, 151, 186, 239, 184, 143, 32, 33})
	actual := GetMessage()

	if !strings.EqualFold(actual, expected) {
		t.Errorf("actual %q expect %q", actual, expected)
	}

	t.Log("OK!")
}
