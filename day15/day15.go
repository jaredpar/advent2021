package day15

import (
	"math"

	"advent2021.com/util"
)

type Cave struct {
	Grid *util.Grid
}

func ParseCave(lines []string) (*Cave, error) {
	grid, err := util.ParseGridFromLines(lines)
	if err != nil {
		return nil, err
	}

	return &Cave{Grid: grid}, nil
}

func Part1(cave *Cave) int {
	grid := cave.Grid
	shortestPath := math.MaxInt
	visited := util.NewGrid(grid.Rows(), grid.Columns())

	var inner func(row, col, cost int)
	inner = func(row, col, cost int) {
		// Once the cost of a given path exceeds the cost of a known
		// path then there is no point going further down it
		if cost > shortestPath {
			return
		}

		if row < 0 ||
			row >= grid.Rows() ||
			col < 0 ||
			col >= grid.Columns() ||
			visited.Value(row, col) == 1 {
			return
		}

		if row != 0 || col != 0 {
			cost += grid.Value(row, col)
		}

		if row+1 == grid.Rows() && col+1 == grid.Columns() {
			shortestPath = util.Min(shortestPath, cost)
			return
		}

		visited.SetValue(row, col, 1)
		defer visited.SetValue(row, col, 0)
		inner(row-1, col, cost)
		inner(row, col+1, cost)
		inner(row+1, col, cost)
		inner(row, col-1, cost)
	}

	inner(0, 0, 0)
	return shortestPath

}
