package main

import (
	"fmt"

	"advent2021.com/day04"
)

func main() {
	result := day04.GetLastScoreFromFile("input.txt", 5)
	fmt.Printf("%d value\n", result)
}
