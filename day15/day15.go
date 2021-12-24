package day15

import (
	"sort"
	"strings"

	"advent2021.com/util"
)

type Cave struct {
	Grid *util.Grid
}

func gridToString(grid *util.Grid) string {
	var sb strings.Builder
	for r := 0; r < grid.Rows(); r++ {
		for c := 0; c < grid.Columns(); c++ {
			digit := grid.Value(r, c)
			sb.WriteRune(util.DigitToRune(digit))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (cave *Cave) String() string {
	var sb strings.Builder
	grid := cave.Grid
	for r := 0; r < grid.Rows(); r++ {
		for c := 0; c < grid.Columns(); c++ {
			digit := grid.Value(r, c)
			sb.WriteRune(util.DigitToRune(digit))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func ParseCave(lines []string) (*Cave, error) {
	grid, err := util.ParseGridFromLines(lines)
	if err != nil {
		return nil, err
	}

	return &Cave{Grid: grid}, nil
}

func shortestPath(grid *util.Grid) int {
	startIndex := grid.Index(0, 0)
	endIndex := grid.Index(grid.Rows()-1, grid.Columns()-1)
	distances := make(map[int]int)
	distances[startIndex] = 0

	doneMap := make(map[int]bool)
	toVisit := []int{startIndex}
	getAdjacentIndexes := func(row, col int) []int {
		indexes := make([]int, 0, 4)
		add := func(row, col int) {
			if row < 0 || row >= grid.Rows() || col < 0 || col >= grid.Columns() {
				return
			}

			indexes = append(indexes, grid.Index(row, col))
		}

		add(row-1, col)
		add(row, col+1)
		add(row+1, col)
		add(row, col-1)
		return indexes
	}

	for {
		index := toVisit[0]
		row, col := grid.RowColumn(index)
		toVisit = toVisit[1:]
		doneMap[index] = true

		distance := distances[index]
		util.Assert(distance >= 0)
		if index == endIndex {
			return distance
		}

		needSort := false
		for _, adjIndex := range getAdjacentIndexes(row, col) {
			if _, done := doneMap[adjIndex]; done {
				continue
			}

			adjRow, adjCol := grid.RowColumn(adjIndex)
			newDistance := grid.Value(adjRow, adjCol) + distance
			oldDistance, present := distances[adjIndex]
			if !present || newDistance < oldDistance {
				if !present {
					toVisit = append(toVisit, adjIndex)
				}

				distances[adjIndex] = newDistance
				needSort = true
			}
		}

		if needSort {
			sort.Slice(toVisit, func(i, j int) bool {
				return distances[toVisit[i]] < distances[toVisit[j]]
			})
		}
	}
}

func Part1(cave *Cave) int {
	return shortestPath(cave.Grid)
}

func Part2(cave *Cave) int {
	oldGrid := cave.Grid
	newGrid := util.NewGrid(oldGrid.Rows()*5, oldGrid.Columns()*5)

	fillSection := func(row, col, inc int) {
		for r := 0; r < oldGrid.Rows(); r++ {
			for c := 0; c < oldGrid.Columns(); c++ {
				value := oldGrid.Value(r, c) + inc
				if value > 9 {
					value -= 9
				}

				newGrid.SetValue(row+r, col+c, value)
			}
		}
	}

	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			fillSection(r*oldGrid.Rows(), c*oldGrid.Columns(), r+c)
		}
	}

	return shortestPath(newGrid)
}
