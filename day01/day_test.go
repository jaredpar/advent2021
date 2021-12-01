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

func TestCountIncreasesWindow(t *testing.T) {
	count := CountIncreasesWindow()
	if count != 1653 {
		t.Error("Incorrect count")
	}

}
