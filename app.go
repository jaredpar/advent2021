package main

import (
	"embed"
	"fmt"

	"advent2021.com/day11"
	"advent2021.com/util"
)

//go:embed day07/*.txt
//go:embed day08/*.txt
//go:embed day09/*.txt
//go:embed day10/*.txt
//go:embed day11/*.txt
var f embed.FS

func main() {

	lines, err := util.ReadLines(f, "day11/input.txt")
	if err != nil {
		panic(err)
	}

	c, err := day11.ParseCavern(lines)
	if err != nil {
		panic(err)
	}

	steps := 0
	for {
		steps++
		flashes := c.RunStep()
		if flashes == c.Count() {
			fmt.Printf("%d\n", steps)
			break
		}
	}
}
