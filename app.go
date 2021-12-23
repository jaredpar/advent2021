package main

import (
	"embed"
	"fmt"

	"advent2021.com/day13"
	"advent2021.com/util"
)

//go:embed day07/*.txt
//go:embed day08/*.txt
//go:embed day09/*.txt
//go:embed day10/*.txt
//go:embed day11/*.txt
//go:embed day12/*.txt
//go:embed day13/*.txt
var f embed.FS

func main() {

	lines, err := util.ReadLines(f, "day13/input.txt")
	if err != nil {
		panic(err)
	}

	m, err := day13.ParseManual(lines)
	if err != nil {
		panic(err)
	}

	m.RunFolds()
	fmt.Println(m.Paper.String())
	fmt.Printf("%d marks\n", m.Paper.CountMarks())
}
