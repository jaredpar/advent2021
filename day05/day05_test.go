package day05

import (
	"testing"

	"advent2021.com/testUtil"
)

func TestOverlapCountOnSample(t *testing.T) {
	d, err := readDiagram("sample.txt")
	testUtil.AssertNotError(t, err)
	testUtil.AssertEqualInt(t, 12, d.board.getOverlapCount())
}

func TestOverlapCountOnInput(t *testing.T) {
	d, err := readDiagram("input.txt")
	testUtil.AssertNotError(t, err)
	testUtil.AssertEqualInt(t, 19258, d.board.getOverlapCount())
}
