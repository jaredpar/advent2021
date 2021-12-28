package day17

import (
	"errors"
	"math"
	"regexp"
	"strconv"

	"advent2021.com/util"
)

type TargetArea struct {
	MinX, MaxX, MinY, MaxY int
}

func NewTargetArea(minX, maxX, minY, maxY int) *TargetArea {
	return &TargetArea{MinX: minX, MaxX: maxX, MinY: minY, MaxY: maxY}
}

func ParseTargetArea(line string) (*TargetArea, error) {
	re := regexp.MustCompile(`target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)`)
	values := re.FindStringSubmatch(line)
	if values == nil {
		return nil, errors.New("invalid line")
	}

	values = values[1:]
	nums := make([]int, 4)
	for i := 0; i < 4; i++ {
		num, err := strconv.Atoi(values[i])
		if err != nil {
			return nil, err
		}

		nums[i] = num
	}

	return NewTargetArea(nums[0], nums[1], nums[2], nums[3]), nil
}

func inRange(val, min, max int) bool {
	if min >= 0 && max >= 0 {
		return val >= min && val <= max
	} else if min >= 0 {
		panic("weird")
	} else {
		if val >= 0 {
			return false
		}

		return val >= min && val <= max
	}
}

func (ta *TargetArea) InTargetAreaX(x int) bool {
	return inRange(x, ta.MinX, ta.MaxX)
}

func (ta *TargetArea) InTargetAreaY(y int) bool {
	return inRange(y, ta.MinY, ta.MaxY)
}

func (ta *TargetArea) InTargetArea(x, y int) bool {
	return ta.InTargetAreaX(x) && ta.InTargetAreaY(y)
}

// Given an initial starting position of (x, y) will it hit the target area?
func (ta *TargetArea) IsHit(x, y int) bool {
	curX := 0
	curY := 0

	for {
		curX += x
		curY += y
		// fmt.Printf("(%d, %d) -> x = %d y = %d\n", curX, curY, x, y)
		if ta.InTargetArea(curX, curY) {
			return true
		}

		if x == 0 && (curX < ta.MinX || curX > ta.MaxX) {
			return false
		}

		if y < 0 && curY < ta.MinY {
			return false
		}

		if x > 0 {
			x--
		} else if x < 0 {
			x++
		}

		y--
	}
}

// Get the set of X velocities that will eventually cross into the target space
func (ta *TargetArea) GetValidX() []int {
	values := make([]int, 0)

	for x := 1; x < ta.MinX; x++ {
		pos := x
		step := x - 1
		for {
			if ta.InTargetAreaX(pos) {
				values = append(values, x)
				break
			}

			if step == 0 || pos > ta.MaxX {
				break
			}

			pos += step
			step--
		}
	}

	for x := ta.MinX; x <= ta.MaxX; x++ {
		values = append(values, x)
	}

	return values
}

func sumRange(max int) int {
	val := 0
	for true {
		val += max
		max--
		if max == 0 {
			break
		}
	}

	return val
}

func PartBoth(line string) (maxY int, distinct int) {
	ta, err := ParseTargetArea(line)
	if err != nil {
		panic(err)
	}

	maxY = math.MinInt
	distinct = 0
	for _, x := range ta.GetValidX() {
		// This is a really brute force method. There is a way to constrain the set
		// of possible 'y' values that I am missing
		for y := util.Min(0, ta.MinY); y < 2000; y++ {
			if ta.IsHit(x, y) {
				//fmt.Printf("(%d, %d)\n", x, y)
				maxY = util.Max(maxY, y)
				distinct++
			}
		}
	}

	return maxY, distinct
}

func Part1(line string) int {
	maxY, _ := PartBoth(line)
	return sumRange(maxY)
}

func Part2(line string) int {
	_, distinct := PartBoth(line)
	return distinct
}
