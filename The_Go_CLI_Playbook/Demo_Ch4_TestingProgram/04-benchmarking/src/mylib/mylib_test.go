package mylib

import (
	"testing"
)

func TestAdder(t *testing.T) {
	if adder(2, 5) != 7 {
		t.Fail()
	}
}

func BenchmarkAdder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		adder(5, 7)
	}
}
