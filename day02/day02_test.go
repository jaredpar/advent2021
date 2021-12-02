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

func TestCalcPart1(t *testing.T) {
	value := CalcPart1()
	if value != 1670340 {
		t.Error(fmt.Sprintf("Value %d is incorrect", value))
	}
}

func TestCalculateWithAimSample(t *testing.T) {
	position, depth, _ := calcPositionWithAim("sample.txt")
	result := position * depth
	if result != 900 {
		t.Error(fmt.Sprintf("Value %d is incorrect", result))
	}
}

func TestCalcPart2(t *testing.T) {
	value := CalcPart2()
	if value != 1954293920 {
		t.Error(fmt.Sprintf("Value %d is incorrect", value))
	}
}
