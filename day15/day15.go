package day15

import (
	"sort"

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
