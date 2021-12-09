package testUtil

import (
	"testing"
)

func AssertEqualInt(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}

func AssertNotError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
