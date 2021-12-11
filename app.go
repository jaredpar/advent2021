package main

import (
	"embed"
	"fmt"

	"advent2021.com/day08"
	"advent2021.com/util"
)

//go:embed day07/*.txt
//go:embed day08/*.txt
var f embed.FS

func main() {

	lines, err := util.ReadLines(f, "day08/input.txt")
	if err != nil {
		panic(err)
	}

	puzzle, err := day08.ParsePuzzle(lines)
	if err != nil {
		panic(err)
	}

	result, err := puzzle.Solve()
	if err != nil {
		panic(err)
	}

	fmt.Printf("result is %d", result)
}
