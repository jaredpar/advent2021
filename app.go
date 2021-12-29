package main

import (
	"embed"
	"fmt"

	"advent2021.com/day15"
	"advent2021.com/day16"
	"advent2021.com/day17"
	"advent2021.com/day18"
	"advent2021.com/util"
)

//go:embed **/*.txt
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

func runDay16Part1() {
	lines, err := util.ReadLines(f, "day16/input.txt")
	if err != nil {
		panic(err)
	}

	sum, err := day16.Part1(lines[0])
	if err != nil {
		panic(err)
	}

	fmt.Printf("version sum: %d\n", sum)
}

func runDay16Part2() {
	lines, err := util.ReadLines(f, "day16/input.txt")
	if err != nil {
		panic(err)
	}

	result, err := day16.Part2(lines[0])
	if err != nil {
		panic(err)
	}

	fmt.Printf("result: %d\n", result)
}

func runDay17Part1() {
	lines := util.MustReadLines(f, "day17/input.txt")
	result := day17.Part1(lines[0])
	fmt.Printf("result: %d\n", result)
}

func runDay17Part2() {
	lines := util.MustReadLines(f, "day17/input.txt")
	result := day17.Part2(lines[0])
	fmt.Printf("result: %d\n", result)
}

func runDay18Part1() {
	lines := util.MustReadLines(f, "day18/input.txt")
	result := day18.Part1(lines)
	fmt.Printf("result: %d\n", result)
}

func runDay18Part2() {
	lines := util.MustReadLines(f, "day18/input.txt")
	result := day18.Part2(lines)
	fmt.Printf("result: %d\n", result)
}

func main() {
	runDay18Part2()
}
