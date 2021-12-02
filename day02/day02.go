package day02

import (
	"embed"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"advent2021.com/util"
)

//go:embed *.txt
var f embed.FS

type direction int
type movement struct {
	direction direction
	value     int
}

const (
	DirectionUp direction = iota
	DirectionDown
	DirectionForward
)

func parseDirection(value string) (direction, error) {
	switch value {
	case "up":
		return DirectionUp, nil
	case "down":
		return DirectionDown, nil
	case "forward":
		return DirectionForward, nil
	default:
		return DirectionUp, nil
	}
}

func readMovements(fileName string) ([]movement, error) {
	lines, err := util.ReadLines(f, fileName)
	if err != nil {
		return nil, err
	}

	movements := make([]movement, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			message := fmt.Sprintf("Input misformed: %s", line)
			return nil, errors.New(message)
		}

		direction, err := parseDirection(parts[0])
		if err != nil {
			message := fmt.Sprintf("Invalid direction: %s", parts[0])
			return nil, errors.New(message)
		}

		value, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		movement := movement{direction: direction, value: value}
		movements = append(movements, movement)
	}

	return movements, nil
}

func calcPosition(fileName string) int {
	movements, err := readMovements(fileName)
	if err != nil {
		panic(err)
	}

	position := 0
	depth := 0

	for _, movement := range movements {
		switch movement.direction {
		case DirectionUp:
			depth -= movement.value
		case DirectionDown:
			depth += movement.value
		case DirectionForward:
			position += movement.value
		default:
			panic("bad enum")
		}
	}

	return position * depth
}

func calcPositionWithAim(fileName string) (position int, depth int, aim int) {
	movements, err := readMovements(fileName)
	if err != nil {
		panic(err)
	}

	for _, movement := range movements {
		switch movement.direction {
		case DirectionUp:
			aim -= movement.value
		case DirectionDown:
			aim += movement.value
		case DirectionForward:
			position += movement.value
			depth += aim * movement.value
		default:
			panic("bad enum")
		}
	}

	return
}

func CalcPart1() int {
	return calcPosition("input.txt")
}

func CalcPart2() int {
	position, depth, _ := calcPositionWithAim("input.txt")
	return position * depth
}
