package day09

import (
	"errors"

	"advent2021.com/util"
)

type FloorMap struct {
	values       []int
	columnLength int
}

func ParseFloorMap(lines []string) (*FloorMap, error) {
	columnLength := -1
	size := 0
	for _, line := range lines {
		size += len(line)
		if columnLength == -1 {
			columnLength = len(line)
		} else if columnLength != len(line) {
			return nil, errors.New("invalid row length")
		}
	}

	if size == 0 {
		return nil, errors.New("no data")
	}

	values := make([]int, size)
	i := 0
	for _, line := range lines {
		for _, r := range line {
			value, err := util.RuneToInt(r)
			if err != nil {
				return nil, err
			}

			values[i] = value
			i++
		}
	}

	return &FloorMap{values: values, columnLength: columnLength}, nil
}

func convertToIndex(row, column, columnLength int) int {
	return (row * columnLength) + column
}

func (f *FloorMap) Rows() int {
	return len(f.values) / f.columnLength
}

func (f *FloorMap) Columns() int {
	return f.columnLength
}

func (f *FloorMap) Value(row, column int) int {
	index := convertToIndex(row, column, f.columnLength)
	return f.values[index]
}

func Part1(f *FloorMap) int {
	isHigherThan := func(value, row, column int) bool {
		if row < 0 || row >= f.Rows() || column < 0 || column >= f.Columns() {
			return true
		}

		return f.Value(row, column) > value
	}

	isLowPoint := func(row, column int) bool {
		value := f.Value(row, column)
		return isHigherThan(value, row-1, column) &&
			isHigherThan(value, row+1, column) &&
			isHigherThan(value, row, column+1) &&
			isHigherThan(value, row, column-1)
	}

	sum := 0
	for r := 0; r < f.Rows(); r++ {
		for c := 0; c < f.Columns(); c++ {
			if isLowPoint(r, c) {
				risk := f.Value(r, c) + 1
				sum += risk
			}
		}
	}

	return sum
}

func Part2(f *FloorMap) int {
	found := make([]bool, len(f.values))
	var calcBasinSize func(row, column int) int
	calcBasinSize = func(row, column int) int {
		if row < 0 || row >= f.Rows() || column < 0 || column >= f.Columns() {
			return 0
		}

		index := convertToIndex(row, column, f.columnLength)
		if found[index] {
			return 0
		}

		found[index] = true
		value := f.values[index]
		if value == 9 {
			return 0
		}

		return 1 +
			calcBasinSize(row-1, column) +
			calcBasinSize(row+1, column) +
			calcBasinSize(row, column+1) +
			calcBasinSize(row, column-1)
	}

	maxes := make([]int, 3)
	updateMaxes := func(size int) {
		if size == 0 {
			return
		}

		for i := 0; i < len(maxes); i++ {
			if size > maxes[i] {
				temp := maxes[i]
				maxes[i] = size
				size = temp
			}
		}
	}

	for r := 0; r < f.Rows(); r++ {
		for c := 0; c < f.Columns(); c++ {
			size := calcBasinSize(r, c)
			updateMaxes(size)
		}
	}

	result := 1
	for _, size := range maxes {
		result *= size
	}

	return result
}
