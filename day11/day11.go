package day11

import (
	"strings"

	"advent2021.com/util"
)

type Cavern struct {
	Grid *util.Grid
}

func ParseCavern(lines []string) (*Cavern, error) {
	grid, err := util.ParseGridFromLines(lines)
	if err != nil {
		return nil, err
	}
	return &Cavern{Grid: grid}, nil
}

// Run a single step and return the number of flashes that occurred
func (cavern *Cavern) RunStep() int {
	grid := cavern.Grid
	count := 0

	// Flash the current octopus and return number of flashes it created
	var flash func(index, row, column int)
	flash = func(index, row, column int) {
		impl := func(row, column int) {
			if row >= 0 && row < grid.Rows() && column >= 0 && column < grid.Columns() {
				index := grid.Index(row, column)
				switch grid.Values[index] {
				case -1:
					// ignore already flashed
				case 9:
					// indirect flash
					flash(index, row, column)
				case 10:
					// normal flash, let the core loop handle
				default:
					grid.Values[index]++
				}
			}
		}

		count++
		grid.Values[index] = -1
		impl(row-1, column)
		impl(row-1, column+1)
		impl(row, column+1)
		impl(row+1, column+1)
		impl(row+1, column)
		impl(row+1, column-1)
		impl(row, column-1)
		impl(row-1, column-1)
	}

	for i := 0; i < len(grid.Values); i++ {
		grid.Values[i]++
	}

	for r := 0; r < grid.Rows(); r++ {
		for c := 0; c < grid.Columns(); c++ {
			index := grid.Index(r, c)
			if grid.Values[index] == 10 {
				flash(index, r, c)
			}
		}
	}

	for i := 0; i < len(grid.Values); i++ {
		if grid.Values[i] >= 10 || grid.Values[i] == -1 {
			grid.Values[i] = 0
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
	grid := cavern.Grid
	var sb strings.Builder
	for r := 0; r < grid.Rows(); r++ {
		for c := 0; c < grid.Columns(); c++ {
			sb.WriteRune(util.DigitToRune(grid.Value(r, c)))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
