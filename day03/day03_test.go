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

func TestGetLifeSupportRatesSample(t *testing.T) {
	result := GetLifeSupportRatesValue("sample.txt")

	if result != 230 {
		t.Errorf("Expected 230 but got %d", result)
	}

}
