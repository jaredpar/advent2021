package day03

import (
	"testing"
)

func TestGetRatesSample(t *testing.T) {
	result := GetRatesValue("sample.txt")

	if result != 198 {
		t.Errorf("Expected 198 but got %d", result)
	}

}
