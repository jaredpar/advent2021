package day10

import (
	"fmt"
	"sort"
	"strings"
)

type stack struct {
	expected rune
	prev     *stack
}

func ParseLine(line string) (rune, string) {
	var s *stack = nil

	push := func(expected rune) {
		s = &stack{expected: expected, prev: s}
	}

	stringIt := func() string {
		if s == nil {
			return ""
		}

		var sb strings.Builder
		for s != nil {
			sb.WriteRune(s.expected)
			s = s.prev
		}
		return sb.String()
	}

	for _, cur := range line {
		switch cur {
		case '{':
			push('}')
		case '<':
			push('>')
		case '[':
			push(']')
		case '(':
			push(')')
		default:
			if s != nil && s.expected != cur {
				return cur, stringIt()
			}

			s = s.prev
		}
	}

	return 0, stringIt()
}

func Part1(lines []string) int {
	score := 0
	for _, line := range lines {
		cur, _ := ParseLine(line)
		if cur == 0 {
			continue
		}

		switch cur {
		case ')':
			score += 3
		case ']':
			score += 57
		case '}':
			score += 1197
		case '>':
			score += 25137
		}
	}

	return score
}

func Part2(lines []string) int {
	scores := make([]int, 0)
	for _, line := range lines {
		score := 0
		add := func(value int) {
			score *= 5
			score += value
		}

		found, expected := ParseLine(line)
		if found != 0 {
			continue
		}

		for _, cur := range expected {
			switch cur {
			case ')':
				add(1)
			case ']':
				add(2)
			case '}':
				add(3)
			case '>':
				add(4)
			default:
				panic(fmt.Sprintf("unexpected character %s", string(cur)))
			}
		}

		if score != 0 {
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	return scores[(len(scores)-1)/2]
}
