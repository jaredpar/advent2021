package day06

import (
	"fmt"
	"strings"
)

type Fish int
type School [9]int

func (s *School) String() string {
	var sb strings.Builder
	for i, count := range *s {
		fmt.Fprintf(&sb, "%d (%d) ", i, count)
	}

	return sb.String()
}

func (s *School) Advance(days int) {
	for ; days > 0; days-- {
		spawned := 0
		for i, count := range *s {
			if i == 0 {
				spawned = count
			} else {
				(*s)[i-1] = count
			}
		}

		(*s)[6] += spawned
		(*s)[8] = spawned
	}
}

func (s School) Count() int {
	count := 0
	for _, c := range s {
		count += c
	}

	return count
}

func NewSchool(values []Fish) School {
	var s School
	for _, f := range values {
		s[int(f)]++
	}
	return s
}
