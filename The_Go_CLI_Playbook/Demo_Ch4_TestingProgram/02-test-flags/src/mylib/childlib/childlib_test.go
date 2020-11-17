package childlib

import (
	"testing"
)

func Test_MoreBasics(t *testing.T) {
	t.Parallel()
	if 10-5 != 5 {
		t.Error("Failed to subtract correctly")
	}
}
