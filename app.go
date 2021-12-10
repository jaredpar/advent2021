package main

import (
	"embed"
	"fmt"

	"advent2021.com/day07"
	"advent2021.com/util"
)

//go:embed day07/*.txt
var f embed.FS

func main() {

	line, err := util.ReadAsSingleLine(f, "day07/input.txt")
	if err != nil {
		panic(err)
	}

	values, err := util.ParseCommaSepInt(line)
	if err != nil {
		panic(err)
	}

	s := day07.NewSwarm(day07.ConvertToCrabs(values))
	position, fuel := s.GetAlignmentEx()
	fmt.Printf("position %d, fuel %d\n", position, fuel)
}
