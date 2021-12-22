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

//go:embed input.txt
var inputText string

func TestSamplePart1(t *testing.T) {
	run := func(text string, expected int) {
		lines := util.ParseLines(text)
		cs, err := ParseCaveSystem(lines)
		if err != nil {
			t.Fatal(err)
		}

		paths := Part1(cs)
		testUtil.AssertEqualInt(t, expected, len(paths))
	}
	run(sampleText, 10)
	run(sample2Text, 19)
	run(inputText, 4707)
}

func TestSamplePart2(t *testing.T) {
	run := func(text string, expected int) {
		lines := util.ParseLines(text)
		cs, err := ParseCaveSystem(lines)
		if err != nil {
			t.Fatal(err)
		}

		paths := Part2(cs)
		testUtil.AssertEqualInt(t, expected, len(paths))
	}
	run(sampleText, 36)
	run(sample2Text, 103)
}
