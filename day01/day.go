package day01

import (
	"embed"

	"advent2021.com/util"
)

//go:embed *.txt
var f embed.FS

func CountIncreases() int {
	lines, err := util.ReadLinesAsInt(f, "input.txt")
	if err != nil || len(lines) == 0 {
		panic("invalid input")
	}

	count := 0
	previous := lines[0]
	lines = lines[1:]
	for _, current := range lines {
		if current > previous {
			count++
		}

		previous = current
	}

	return count
}

func CountIncreasesWindow() int {
	lines, err := util.ReadLinesAsInt(f, "input.txt")
	if err != nil {
		panic("invalid input")
	}

	sumRange := func(index int) int {
		return lines[index] + lines[index+1] + lines[index+2]
	}

	count := 0
	for i := 0; i+3 < len(lines); i++ {
		if sumRange(i+1) > sumRange(i) {
			count++
		}
	}

	return count
}
