package day06

import (
	"testing"

	"advent2021.com/testUtil"
)

func TestSample(t *testing.T) {
	s := NewSchool([]Fish{3, 4, 3, 1, 2})
	s.Advance(80)
	testUtil.AssertEqualInt(t, 5934, s.Count())
}
