package day17

import (
	"testing"

	"advent2021.com/testUtil"
)

func mustParse(t *testing.T, line string) *TargetArea {
	ta, err := ParseTargetArea(line)
	if err != nil {
		t.Fatal(err)
	}

	return ta
}

func TestParseTargetArea(t *testing.T) {
	testCore := func(line string, minX, maxX, minY, maxY int) {
		ta, err := ParseTargetArea(line)
		if err != nil {
			t.Errorf("cannot parse: %s", line)
		}

		testUtil.AssertEqualInt(t, minX, ta.MinX)
		testUtil.AssertEqualInt(t, maxX, ta.MaxX)
		testUtil.AssertEqualInt(t, minY, ta.MinY)
		testUtil.AssertEqualInt(t, maxY, ta.MaxY)
	}

	testCore("target area: x=153..199, y=-114..-75", 153, 199, -114, -75)
	testCore("target area: x=20..30, y=-10..-5", 20, 30, -10, -5)
}

func TestSimpleHit(t *testing.T) {
	ta := mustParse(t, "target area: x=20..30, y=-10..-5")
	testUtil.Assert(t, ta.IsHit(7, 2))
	testUtil.Assert(t, ta.IsHit(6, 3))
	testUtil.Assert(t, !ta.IsHit(17, -4))

}
