package day02

import (
	"fmt"
	"testing"
)

func TestCalculatePositionSample(t *testing.T) {
	value := calcPosition("sample.txt")
	if value != 150 {
		t.Error(fmt.Sprintf("Value %d is incorrect", value))
	}
}
