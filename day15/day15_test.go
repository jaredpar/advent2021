package day15

import (
	_ "embed"
	"testing"

	"advent2021.com/testUtil"
	"advent2021.com/util"
)

//go:embed sample.txt
var sampleText string

//go:embed input.txt
var inputText string

func TestPart1(t *testing.T) {
	run := func(text string, expected int) {
		lines := util.ParseLines(text)
		cave, err := ParseCave(lines)
		if err != nil {
			t.Fatal(err)
		}

		cost := Part1(cave)
		testUtil.AssertEqualInt(t, expected, cost)
	}

	run(sampleText, 40)
	run(inputText, 589)
}

func TestPart2(t *testing.T) {
	run := func(text string, expected int) {
		lines := util.ParseLines(text)
		cave, err := ParseCave(lines)
		if err != nil {
			t.Fatal(err)
		}

		cost := Part2(cave)
		testUtil.AssertEqualInt(t, expected, cost)
	}

	run(sampleText, 315)
	run(inputText, 2885)
}
