package day14

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"advent2021.com/util"
)

type Data struct {
	Template string
	Rules    map[util.RunePair]string
}

func ParseData(lines []string) (*Data, error) {
	if len(lines) < 3 {
		return nil, errors.New("bad data")
	}

	template := lines[0]
	rules := make(map[util.RunePair]string)
	lines = lines[2:]

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 3 || parts[1] != "->" {
			return nil, fmt.Errorf("bad input line: %s", line)
		}

		runes := []rune(parts[0])
		if len(runes) != 2 {
			return nil, fmt.Errorf("bad key: %s", parts[0])
		}

		pair := util.NewRunePair(runes[0], runes[1])
		rules[pair] = parts[2]
	}

	return &Data{Template: template, Rules: rules}, nil
}

func (d *Data) Run(steps int) string {
	line := []rune(d.Template)
	runOne := func() {
		var sb strings.Builder
		for i := 0; i < len(line); i++ {

			if i+1 == len(line) {
				sb.WriteRune(line[i])
			} else {
				sb.WriteRune(line[i])
				pair := util.NewRunePair(line[i], line[i+1])
				if insert, ok := d.Rules[pair]; ok {
					sb.WriteString(insert)
				}
			}
		}

		line = []rune(sb.String())
	}

	for i := 0; i < steps; i++ {
		runOne()
	}

	return string(line)
}

func Part1(d *Data) int {
	result := d.Run(10)
	countMap := make(map[rune]int)

	for _, r := range result {
		count, ok := countMap[r]
		if !ok {
			count = 0
		}

		count++
		countMap[r] = count
	}

	max := 0
	min := math.MaxInt
	for _, v := range countMap {
		max = util.Max(max, v)
		min = util.Min(min, v)
	}

	return max - min
}

/*
func Part2(d *Data) int {
	countMap := make(map[rune]int)
	run := func(left, right rune)

	for _, r := range result {
		count, ok := countMap[r]
		if !ok {
			count = 0
		}

		count++
		countMap[r] = count
	}

	max := 0
	min := math.MaxInt
	for _, v := range countMap {
		max = util.Max(max, v)
		min = util.Min(min, v)
	}

	return max - min
}

*/
