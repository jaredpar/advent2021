package main

import (
	"embed"
	"fmt"

	"advent2021.com/day09"
	"advent2021.com/util"
)

//go:embed day07/*.txt
//go:embed day08/*.txt
//go:embed day09/*.txt
var f embed.FS

func main() {

	lines, err := util.ReadLines(f, "day09/input.txt")
	if err != nil {
		panic(err)
	}

	f, err := day09.ParseFloorMap(lines)
	if err != nil {
		panic(err)
	}

	sum := day09.Part2(f)
	fmt.Printf("result is %d", sum)
}
