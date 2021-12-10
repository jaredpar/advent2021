package day07

import (
	"math"

	"advent2021.com/util"
)

type Crab int
type Swarm []Crab

func ConvertToCrabs(values []int) []Crab {
	crabs := make([]Crab, len(values))
	for i, v := range values {
		crabs[i] = Crab(v)
	}
	return crabs
}

func NewSwarm(crabs []Crab) Swarm {
	return Swarm(crabs)
}

// This is effectively bubble sort but should work for the input size
// here. Do need to look up more effective algorithm here
func (s Swarm) GetAlignment() (position, fuel int) {
	min := math.MaxInt
	max := math.MinInt

	for _, c := range s {
		i := int(c)
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}

	position = 0
	fuel = math.MaxInt
	for p := min; p <= max; p++ {
		current := 0
		for _, c := range s {
			i := int(c)
			current += util.Abs(i - p)
		}

		if current < fuel {
			fuel = current
			position = p
		}
	}

	return
}

func (s Swarm) GetAlignmentEx() (position, fuel int) {
	min := math.MaxInt
	max := math.MinInt

	for _, c := range s {
		i := int(c)
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}

	position = 0
	fuel = math.MaxInt

	// The fuel cost here is just an arithmic series
	getFuelCost := func(len int) int {
		if len == 1 {
			return 1
		}

		return (len * (1 + len)) / 2
	}

	for p := min; p <= max; p++ {
		current := 0
		for _, c := range s {
			i := int(c)
			current += getFuelCost(util.Abs(i - p))
		}

		if current < fuel {
			fuel = current
			position = p
		}
	}

	return
}
