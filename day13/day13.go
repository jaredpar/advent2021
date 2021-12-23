package day13

import (
	"fmt"
	"strconv"
	"strings"

	"advent2021.com/util"
)

type Paper struct {
	Grid *util.Grid
}

func NewPaper(g *util.Grid) *Paper {
	return &Paper{Grid: g}
}

func (p *Paper) Fold(isRow bool, value int) {
	g := p.Grid
	if isRow {
		limit := util.Min(value, g.Rows()-(value+1))
		for i := 1; i <= limit; i++ {
			for c := 0; c < g.Columns(); c++ {
				if g.Value(i+value, c) == 1 {
					g.SetValue(value-i, c, 1)
				}
			}
		}
		g.Resize(value, g.Columns())
	} else {
		limit := util.Min(value, g.Columns()-(value+1))
		for r := 0; r < g.Rows(); r++ {
			for i := 1; i <= limit; i++ {
				if g.Value(r, i+value) == 1 {
					g.SetValue(r, value-i, 1)
				}
			}
		}
		g.Resize(g.Rows(), value)
	}
}

func (p *Paper) CountMarks() int {
	count := 0
	for _, v := range p.Grid.Values {
		if v == 1 {
			count++
		}
	}

	return count
}

func (p *Paper) String() string {
	var sb strings.Builder
	for r := 0; r < p.Grid.Rows(); r++ {
		for c := 0; c < p.Grid.Columns(); c++ {
			if p.Grid.Value(r, c) == 1 {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

type Fold struct {
	IsRow bool
	Value int
}

func NewFold(isRow bool, value int) Fold {
	return Fold{IsRow: isRow, Value: value}
}

type Manual struct {
	Paper *Paper
	Folds []Fold
}

func ParseManual(lines []string) (*Manual, error) {
	g := util.NewGrid(1, 1)
	foldIndex := 0

	for i, line := range lines {
		if len(line) == 0 {
			foldIndex = i + 1
			break
		}

		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid syntax: %s", line)
		}

		column, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		row, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		g.Resize(util.Max(g.Rows(), row+1), util.Max(g.Columns(), column+1))
		g.SetValue(row, column, 1)
	}

	lines = lines[foldIndex:]
	folds := make([]Fold, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid syntax: %s", line)
		}

		notations := strings.Split(parts[2], "=")
		if len(notations) != 2 {
			return nil, fmt.Errorf("invalid fold notation: %s", notations[2])
		}

		var isRow bool
		switch notations[0] {
		case "x":
			isRow = false
		case "y":
			isRow = true
		default:
			return nil, fmt.Errorf("expected x or y but got %s", notations[0])
		}

		value, err := strconv.Atoi(notations[1])
		if err != nil {
			return nil, err
		}

		folds = append(folds, NewFold(isRow, value))
	}

	return &Manual{Paper: NewPaper(g), Folds: folds}, nil
}
