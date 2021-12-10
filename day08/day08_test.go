package day08

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
	puzzle, err := ParsePuzzle(lines)
	if err != nil {
		t.Error(err)
		return
	}

	testUtil.AssertEqualInt(t, 26, puzzle.GetKnownOutputCount())
}
