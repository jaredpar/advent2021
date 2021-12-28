package day17

import (
	"errors"
	"fmt"
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
	return val >= min && val <= max
}

func (ta *TargetArea) InTargetAreaX(x int) bool {
	return inRange(x, ta.MinX, ta.MaxX)
}

func (ta *TargetArea) InTargetAreaY(x int) bool {
	return inRange(x, ta.MinY, ta.MaxY)
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

		if curX == 0 && (curX < ta.MinX || curX > ta.MaxX) {
			return false
		}

		if curY <= 0 && curY < ta.MinY {
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

func Part1(line string) int {
	ta, err := ParseTargetArea(line)
	if err != nil {
		panic(err)
	}

	getValidX := func() []int {
		values := make([]int, 0)

		x := 0
		foundHit := false
		for {
			if ta.InTargetAreaX(sumRange(x + 1)) {
				x++
				foundHit = true
				values = append(values, x)
			} else if foundHit {
				break
			} else {
				x++
			}
		}

		return values
	}

	maxY := 0
	for _, x := range getValidX() {
		fmt.Printf("trying x: %d\n", x)

		// Try y values until we find the max
		y := 0
		foundHit := false
		for {
			if ta.IsHit(x, y+1) {
				fmt.Printf("(%d, %d)\n", x, y)
				foundHit = true
			} else if foundHit {
				break
			}

			y++
		}

		maxY = util.Max(y, maxY)
	}

	return maxY
}
