package mylib

import (
	"testing"
)

func TestMessageWriter(t *testing.T) {
	if messageWriter("Hello", "World") != "Hello, World" {
		t.Fail()
	}
}
