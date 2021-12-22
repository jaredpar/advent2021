package day11

import (
	"errors"
	"strings"

	"advent2021.com/util"
)

type Cavern struct {
	values       []int
	columnLength int
}

func ParseCavern(lines []string) (*Cavern, error) {
	values := make([]int, 0)
	columnLength := 0
	for _, line := range lines {
		if columnLength == 0 {
			columnLength = len(line)
		} else if columnLength != len(line) {
			return nil, errors.New("invalid column length")
		}

		for _, r := range line {
			value, err := util.RuneToInt(r)
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
	}

	return &Cavern{values: values, columnLength: columnLength}, nil
}

func (c *Cavern) Index(row, column int) int {
	return (row * c.columnLength) + column
}

func (c *Cavern) Rows() int {
	return len(c.values) / c.columnLength
}

func (c *Cavern) Columns() int {
	return c.columnLength
}

func (c *Cavern) Value(row, column int) int {
	index := c.Index(row, column)
	return c.values[index]
}

// Run a single step and return the number of flashes that occurred
func (cavern *Cavern) RunStep() int {
	count := 0

	// Flash the current octopus and return number of flashes it created
	var flash func(index, row, column int)
	flash = func(index, row, column int) {
		impl := func(row, column int) {
			if row >= 0 && row < cavern.Rows() && column >= 0 && column < cavern.Columns() {
				index := cavern.Index(row, column)
				switch cavern.values[index] {
				case -1:
					// ignore already flashed
				case 9:
					// indirect flash
					flash(index, row, column)
				case 10:
					// normal flash, let the core loop handle
				default:
					cavern.values[index]++
				}
			}
		}

		count++
		cavern.values[index] = -1
		impl(row-1, column)
		impl(row-1, column+1)
		impl(row, column+1)
		impl(row+1, column+1)
		impl(row+1, column)
		impl(row+1, column-1)
		impl(row, column-1)
		impl(row-1, column-1)
	}

	for i := 0; i < len(cavern.values); i++ {
		cavern.values[i]++
	}

	for r := 0; r < cavern.Rows(); r++ {
		for c := 0; c < cavern.Columns(); c++ {
			index := cavern.Index(r, c)
			if cavern.values[index] == 10 {
				flash(index, r, c)
			}
		}
	}

	for i := 0; i < len(cavern.values); i++ {
		if cavern.values[i] >= 10 || cavern.values[i] == -1 {
			cavern.values[i] = 0
		}
	}

	return count
}

func (cavern *Cavern) RunSteps(steps int) int {
	count := 0
	for i := 0; i < steps; i++ {
		count += cavern.RunStep()
	}

	return count
}

func (cavern *Cavern) String() string {
	var sb strings.Builder
	for r := 0; r < cavern.Rows(); r++ {
		for c := 0; c < cavern.Columns(); c++ {
			sb.WriteRune(util.DigitToRune(cavern.Value(r, c)))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
