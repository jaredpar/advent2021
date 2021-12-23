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
	Rules    map[util.RunePair]rune
}

func ParseData(lines []string) (*Data, error) {
	if len(lines) < 3 {
		return nil, errors.New("bad data")
	}

	template := lines[0]
	rules := make(map[util.RunePair]rune)
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
		runes = []rune(parts[2])
		if len(runes) != 1 {
			return nil, fmt.Errorf("bad value: %s", parts[2])
		}

		rules[pair] = runes[0]
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
					sb.WriteRune(insert)
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

func dumpRuneMap(runeMap map[rune]int) {
	for r, c := range runeMap {
		fmt.Printf("%s -> %d\n", string(r), c)
	}
	fmt.Println()
}

func Part2(d *Data, steps int) int {
	// Hold the number of times a given pair occurs in the sequence at this point
	pairMap := make(map[util.RunePair]int)
	incrementPair := func(pair util.RunePair, inc int) {
		count, _ := pairMap[pair]
		count += inc
		pairMap[pair] = count
	}

	copyPairMap := func() map[util.RunePair]int {
		m := make(map[util.RunePair]int)
		for k, v := range pairMap {
			m[k] = v
		}
		return m
	}

	// Holds the number of times a given character will appear in the final output
	runeMap := make(map[rune]int)
	runes := []rune(d.Template)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if i+1 < len(runes) {
			pair := util.NewRunePair(r, runes[i+1])
			incrementPair(pair, 1)
		}
		runeCount, _ := runeMap[r]
		runeMap[r] = runeCount + 1
	}

	for i := 0; i < steps; i++ {
		pairMapCopy := copyPairMap()
		for pair, r := range d.Rules {
			count, _ := pairMapCopy[pair]
			pair1 := util.NewRunePair(pair.Left, r)
			incrementPair(pair1, count)

			pair2 := util.NewRunePair(r, pair.Right)
			if pair2 != pair1 {
				incrementPair(pair2, count)
			}
			curCount, _ := pairMap[pair]
			pairMap[pair] = curCount - count

			runeCount, _ := runeMap[r]
			runeMap[r] = runeCount + count
		}
	}

	max := 0
	min := math.MaxInt
	for _, v := range runeMap {
		max = util.Max(max, v)
		min = util.Min(min, v)
	}

	return max - min
}
