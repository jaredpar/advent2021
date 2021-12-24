package util

import (
	"errors"
	"fmt"
)

type Grid struct {
	Values       []int
	ColumnLength int
}

func rowColumn(index, columnLength int) (row, column int) {
	column = index % columnLength
	row = (index - column) / columnLength
	return
}

func index(row, column, columnLength int) int {
	return (row * columnLength) + column
}

func NewGrid(row, column int) *Grid {
	if row == 0 || column == 0 {
		panic("row and column must be above 0")
	}

	values := make([]int, row*column)
	return &Grid{Values: values, ColumnLength: column}
}

func (g *Grid) Count() int {
	return len(g.Values)
}

func (g *Grid) Index(row, column int) int {
	return index(row, column, g.ColumnLength)
}

func (g *Grid) RowColumn(index int) (row, column int) {
	return rowColumn(index, g.ColumnLength)
}

func (g *Grid) Rows() int {
	return len(g.Values) / g.ColumnLength
}

func (g *Grid) Columns() int {
	return g.ColumnLength
}

func (g *Grid) Value(row, column int) int {
	index := g.Index(row, column)
	return g.Values[index]
}

func (g *Grid) SetValue(row, column, value int) {
	index := g.Index(row, column)
	g.Values[index] = value
}

func (g *Grid) SetAll(value int) {
	SetAllInt(g.Values, value)
}

func (g *Grid) Resize(row, column int) {
	if row == g.Rows() && column == g.Columns() {
		return
	}

	oldColumnLength := g.ColumnLength
	oldValues := g.Values

	g.Values = make([]int, row*column)
	g.ColumnLength = column

	for i, v := range oldValues {
		r, c := rowColumn(i, oldColumnLength)
		if r < row && c < column {
			g.SetValue(r, c, v)
		}
	}
}

// Parse out a grid from a series of single digit entries on a
// line like the following
//
// 0123
// 4567
func ParseGridFromLines(lines []string) (*Grid, error) {
	if len(lines) == 0 {
		return nil, errors.New("need at least one line")
	}

	grid := NewGrid(len(lines), len(lines[0]))
	for row, line := range lines {
		if len(line) != grid.Rows() {
			return nil, fmt.Errorf("line has wrong length: %s", line)
		}

		for col, r := range line {
			digit, err := RuneToInt(r)
			if err != nil {
				return nil, err
			}

			grid.SetValue(row, col, digit)
		}
	}

	return grid, nil
}
