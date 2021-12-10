package day07

import (
	_ "embed"
	"testing"

	"advent2021.com/testUtil"
	"advent2021.com/util"
)

//go:embed input.txt
var inputText string

//go:embed sample.txt
var sampleText string

func testFuelCore(t *testing.T, line string, expectedPosition int, expectedFuel int) {
	values, err := util.ParseCommaSepInt(line)
	if err != nil {
		t.Error(err)
		return
	}

	crabs := ConvertToCrabs(values)
	swarm := NewSwarm(crabs)
	position, fuel := swarm.GetAlignment()
	testUtil.AssertEqualInt(t, expectedPosition, position)
	testUtil.AssertEqualInt(t, expectedFuel, fuel)
}

func TestSample(t *testing.T) {
	testFuelCore(t, sampleText, 2, 37)
}

func TestInput(t *testing.T) {
	testFuelCore(t, inputText, 371, 341558)
}
