package day03

import (
	"embed"
	"errors"
	"math"

	"advent2021.com/util"
)

//go:embed *.txt
var f embed.FS

func GetRates(fileName string) (gamma int, epsilon int, err error) {
	lines, err := util.ReadLines(f, fileName)
	if err != nil {
		return 0, 0, err
	}

	if len(lines) == 0 {
		return 0, 0, nil
	}

	// Inefficient for now but will get the job done
	length := len(lines[0])
	counters := make([]int, length)

	for _, line := range lines {
		if len(line) != length {
			return 0, 0, errors.New("Invalid line length")
		}

		for i, c := range line {
			if c == '0' {
				counters[i]--
			} else {
				counters[i]++
			}
		}
	}

	for i, counter := range counters {
		exp := (length - 1) - i
		base := int(math.Pow(2, float64(exp)))
		if counter > 0 {
			// 1 is more common
			gamma += base
		} else {
			// 0 is more common
			epsilon += base
		}
	}

	return
}

func GetRatesValue(fileName string) int {
	gamma, epsilon, err := GetRates(fileName)
	if err != nil {
		panic(err)
	}

	return gamma * epsilon
}
