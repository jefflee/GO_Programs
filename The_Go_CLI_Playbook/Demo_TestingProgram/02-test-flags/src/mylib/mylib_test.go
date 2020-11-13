package mylib

import (
	"testing"
)

func Test_BasicChecks(t *testing.T) {
	t.Parallel()
	t.Run("Go can add", func(t *testing.T) {
		if 1+1 != 2 {
			t.Fail()
		}
	})
	t.Run("Go can concatenate strings", func(t *testing.T) {
		if "Hello, "+"Go" != "Hello, Go" {
			t.Fail()
		}
	})
}
