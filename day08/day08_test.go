package day08

import (
	_ "embed"
	"testing"

	"advent2021.com/testUtil"
	"advent2021.com/util"
)

//go:embed sample.txt
var sampleText string

func TestSample(t *testing.T) {
	lines := util.ParseLines(sampleText)
	puzzle, err := ParsePuzzle(lines)
	if err != nil {
		t.Error(err)
		return
	}

	testUtil.AssertEqualInt(t, 26, puzzle.GetKnownEasyOutputCount())

	result, err := puzzle.Solve()
	if err != nil {
		t.Error(err)
		return
	}

	testUtil.AssertEqualInt(t, 61229, result)
}

func TestLearnMappings(t *testing.T) {
	inputs := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	solved, err := learnMapping(inputs)
	if err != nil {
		t.Error(err)
	}

	assertMap := func(from, to byte) {
		value := solved[from]
		if value != to {
			t.Errorf("expected %c -> %c but found %c", from, to, value)
		}
	}

	assertMap('d', 'a')
	assertMap('e', 'b')
	assertMap('a', 'c')
	assertMap('f', 'd')
	assertMap('g', 'e')
	assertMap('b', 'f')
	assertMap('c', 'g')

	outputs := []string{"cdfeb", "fcadb", "cdfeb", "cdbaf"}
	entry := Entry{Input: inputs, Output: outputs}
	result, err := solveEntry(&entry)
	if err != nil {
		t.Error(err)
	}

	testUtil.AssertEqualInt(t, 5353, result)
}
