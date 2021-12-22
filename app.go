package main

import (
	"embed"
	"fmt"

	"advent2021.com/day12"
	"advent2021.com/util"
)

//go:embed day07/*.txt
//go:embed day08/*.txt
//go:embed day09/*.txt
//go:embed day10/*.txt
//go:embed day11/*.txt
//go:embed day12/*.txt
var f embed.FS

func main() {

	lines, err := util.ReadLines(f, "day12/input.txt")
	if err != nil {
		panic(err)
	}

	cs, err := day12.ParseCaveSystem(lines)
	if err != nil {
		panic(err)
	}

	paths := day12.Part1(cs)
	for _, path := range paths {
		fmt.Println(path)
	}

	fmt.Printf("%d paths", len(paths))
}
