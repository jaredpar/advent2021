package day01

import (
	"testing"
)

func TestCountIncreases(t *testing.T) {
	count := CountIncreases()
	if count != 1624 {
		t.Error("Incorrect count")
	}

}
