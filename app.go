package main

import (
	"embed"
	"fmt"

	"advent2021.com/day15"
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
//go:embed day15/*.txt
var f embed.FS

func runDay15Part1() {
	lines, err := util.ReadLines(f, "day15/input.txt")
	if err != nil {
		panic(err)
	}

	cave, err := day15.ParseCave(lines)
	if err != nil {
		panic(err)
	}

	result := day15.Part1(cave)
	fmt.Printf("%d\n", result)
}

func runDay15Part2() {
	lines, err := util.ReadLines(f, "day15/input.txt")
	if err != nil {
		panic(err)
	}

	cave, err := day15.ParseCave(lines)
	if err != nil {
		panic(err)
	}

	result := day15.Part2(cave)
	fmt.Printf("%d\n", result)
}

func main() {
	runDay15Part2()
}
