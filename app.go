package main

import (
	"embed"
	"fmt"

	"advent2021.com/day14"
	"advent2021.com/util"
)

//go:embed day07/*.txt
//go:embed day08/*.txt
//go:embed day09/*.txt
//go:embed day10/*.txt
//go:embed day11/*.txt
//go:embed day12/*.txt
//go:embed day13/*.txt
//go:embed day14/*.txt
var f embed.FS

func main() {

	lines, err := util.ReadLines(f, "day14/input.txt")
	if err != nil {
		panic(err)
	}

	d, err := day14.ParseData(lines)
	if err != nil {
		panic(err)
	}

	result := day14.Part2(d, 40)
	fmt.Printf("%d\n", result)
}
