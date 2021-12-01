package main

import (
	"fmt"

	"advent2021.com/day01"
)

func assert(cond bool) {
	if !cond {
		panic("failed assert")
	}
}

func verify() {

}

func main() {
	count := day01.CountIncreases()
	fmt.Printf("%d increases\n", count)

}
