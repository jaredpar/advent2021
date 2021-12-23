package util

import (
	"testing"

	"advent2021.com/testUtil"
)

func TestRowColumn(t *testing.T) {
	g := NewGrid(10, 10)
	for r := 0; r < g.Rows(); r++ {
		for c := 0; c < g.Columns(); c++ {
			index := g.Index(r, c)
			actualRow, actualCol := g.RowColumn(index)
			testUtil.AssertEqualInt(t, r, actualRow)
			testUtil.AssertEqualInt(t, c, actualCol)
		}
	}
}

func TestExpand(t *testing.T) {
	g := NewGrid(5, 5)
	for i := 0; i < g.Count(); i++ {
		g.Values[i] = i
	}

	g.Resize(10, 10)
	for r := 0; r < g.Rows(); r++ {
		for c := 0; c < g.Columns(); c++ {
			if r < 5 && c < 5 {
				value := index(r, c, 5)
				testUtil.AssertEqualInt(t, value, g.Value(r, c))
			} else {
				testUtil.AssertEqualInt(t, 0, g.Value(r, c))
			}
		}
	}
}

func TestShrink(t *testing.T) {
	g := NewGrid(5, 5)
	for i := 0; i < g.Count(); i++ {
		g.Values[i] = 1
	}

	g.Resize(3, 3)
	for r := 0; r < g.Rows(); r++ {
		for c := 0; c < g.Columns(); c++ {
			testUtil.AssertEqualInt(t, 1, g.Value(r, c))
		}
	}
}
