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

	count := puzzle.GetKnownOutputCount()
	fmt.Printf("count is %d", count)
}
