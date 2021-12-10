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
