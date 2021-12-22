package day11

import (
	_ "embed"
	"testing"

	"advent2021.com/testUtil"
	"advent2021.com/util"
)

//go:embed sample.txt
var sampleText string

func TestSamplePart1(t *testing.T) {
	lines := util.ParseLines(sampleText)
	c, err := ParseCavern(lines)
	if err != nil {
		t.Fatal(err)
	}

	flashes := c.RunSteps(100)
	testUtil.AssertEqualInt(t, 1656, flashes)
}

func TestSamplePart2(t *testing.T) {
	lines := util.ParseLines(sampleText)
	c, err := ParseCavern(lines)
	if err != nil {
		t.Fatal(err)
	}

	steps := 0
	for {
		steps++
		flashes := c.RunStep()
		if flashes == c.Grid.Count() {
			break
		}
	}
	testUtil.AssertEqualInt(t, 195, steps)
}
