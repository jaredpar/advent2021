package main

import (
	"embed"
	"fmt"

	"advent2021.com/day10"
	"advent2021.com/util"
)

//go:embed day07/*.txt
//go:embed day08/*.txt
//go:embed day09/*.txt
//go:embed day10/*.txt
var f embed.FS

func main() {

	lines, err := util.ReadLines(f, "day10/input.txt")
	if err != nil {
		panic(err)
	}

	score := day10.Part1(lines)
	fmt.Printf("Score is %d\n", score)
}
