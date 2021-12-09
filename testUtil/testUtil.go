package testUtil

import (
	"testing"
)

func AssertEqualInt(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
