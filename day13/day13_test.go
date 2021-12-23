package day13

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

func TestSample(t *testing.T) {
	lines := util.ParseLines(sampleText)
	m, err := ParseManual(lines)
	if err != nil {
		t.Fatal(err)
	}

	if len(m.Folds) != 2 {
		t.Fatalf("expected 2 folds but got %d", len(m.Folds))
	}

	for _, fold := range m.Folds {
		m.Paper.Fold(fold.IsRow, fold.Value)
	}

	testUtil.AssertEqualInt(t, 16, m.Paper.CountMarks())
}

func TestPart1(t *testing.T) {
	lines := util.ParseLines(sampleText)
	m, err := ParseManual(lines)
	if err != nil {
		t.Fatal(err)
	}

	m.Paper.Fold(true, 7)
	testUtil.AssertEqualInt(t, 17, m.Paper.CountMarks())
}
