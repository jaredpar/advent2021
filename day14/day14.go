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
	Rules    map[string]string
}

func ParseData(lines []string) (*Data, error) {
	if len(lines) < 3 {
		return nil, errors.New("bad data")
	}

	template := lines[0]
	rules := make(map[string]string)
	lines = lines[2:]

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 3 || parts[1] != "->" {
			return nil, fmt.Errorf("bad input line: %s", line)
		}

		rules[parts[0]] = parts[2]
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
				if insert, ok := d.Rules[string(line[i:i+2])]; ok {
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
