package day01

import (
	"embed"

	"advent2021.com/util"
)

//go:embed input.txt
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
