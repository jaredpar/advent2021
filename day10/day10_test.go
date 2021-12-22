package day10

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
	score := Part1(lines)
	testUtil.AssertEqualInt(t, 26397, score)
}

func TestSamplePart2(t *testing.T) {
	lines := util.ParseLines(sampleText)
	score := Part2(lines)
	testUtil.AssertEqualInt(t, 288957, score)
}
