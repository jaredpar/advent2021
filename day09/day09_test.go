package day09

import (
	_ "embed"
	"testing"

	"advent2021.com/testUtil"
	"advent2021.com/util"
)

//go:embed sample.txt
var sampleText string

func TestPart1Sample(t *testing.T) {
	lines := util.ParseLines(sampleText)
	f, err := ParseFloorMap(lines)
	if err != nil {
		t.Fatal(err)
	}

	sum := Part1(f)
	testUtil.AssertEqualInt(t, 15, sum)
}

func TestPart2Sample(t *testing.T) {
	lines := util.ParseLines(sampleText)
	f, err := ParseFloorMap(lines)
	if err != nil {
		t.Fatal(err)
	}

	result := Part2(f)
	testUtil.AssertEqualInt(t, 1134, result)
}
