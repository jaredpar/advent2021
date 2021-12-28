package day17

import (
	"errors"
	"regexp"
	"strconv"
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
