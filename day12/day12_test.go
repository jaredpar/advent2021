package day12

import (
	_ "embed"
	"testing"

	"advent2021.com/testUtil"
	"advent2021.com/util"
)

//go:embed sample.txt
var sampleText string

//go:embed sample2.txt
var sample2Text string

func TestSamplePart1(t *testing.T) {
	lines := util.ParseLines(sampleText)
	cs, err := ParseCaveSystem(lines)
	if err != nil {
		t.Fatal(err)
	}

	paths := Part1(cs)
	testUtil.AssertEqualInt(t, 10, len(paths))
}

func TestSample2Part1(t *testing.T) {
	lines := util.ParseLines(sample2Text)
	cs, err := ParseCaveSystem(lines)
	if err != nil {
		t.Fatal(err)
	}

	paths := Part1(cs)
	testUtil.AssertEqualInt(t, 19, len(paths))
}
