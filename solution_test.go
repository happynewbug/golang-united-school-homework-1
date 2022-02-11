package main

import "testing"

func TestGetMessage(t *testing.T) {
	actual := GetMessage()
	expect := "Hello 🗺️!"

	if actual != expect {
		t.Errorf("actual %q expect %q", actual, expect)
	}
}
