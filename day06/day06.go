package day06

import (
	"fmt"
	"strings"

	"advent2021.com/util"
)

type Fish int
type School []Fish

func NewFish(values []int) []Fish {
	fishes := make([]Fish, len(values))
	for i, v := range values {
		fishes[i] = Fish(v)
	}
	return fishes
}

func (s *School) String() string {
	var counts [9]int

	for _, cur := range *s {
		util.Assert(cur <= 8)
		counts[cur]++
	}

	var sb strings.Builder
	for i, count := range counts {
		fmt.Fprintf(&sb, "%d (%d) ", i, count)
	}

	return sb.String()
}

func (s *School) Advance(days int) {
	fishes := []Fish(*s)
	for ; days > 0; days-- {
		spawned := 0
		for i, fish := range fishes {
			if fish == 0 {
				spawned++
				fishes[i] = 6
			} else {
				fishes[i]--
			}
		}

		for i := 0; i < spawned; i++ {
			fishes = append(fishes, 8)
		}
	}

	*s = fishes
}

func (s School) Count() int {
	return len(s)
}

func NewSchool(values []Fish) School {
	return School(values)
}
