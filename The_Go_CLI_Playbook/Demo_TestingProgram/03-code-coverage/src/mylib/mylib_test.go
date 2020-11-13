package mylib

import (
	"testing"
)

func TestAdder(t *testing.T) {
	if adder(2, 5) != 7 {
		t.Fail()
	}
}
