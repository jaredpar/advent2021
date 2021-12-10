package day08

import (
	"errors"

	"advent2021.com/util"
)

type Display int

const (
	DisplayZero Display = iota
	DisplayOne
	DisplayTwo
	DisplayThree
	DisplayFour
	DisplayFive
	DisplaySix
	DisplaySeven
	DisplayEight
	DisplayNine
	DisplayUnknown
)

type Entry struct {
	Input  []string
	Output []string
}

type Puzzle []*Entry

func ParseEntry(line string) (*Entry, error) {
	parts := util.SplitOnWhiteSpace(line)
	if len(parts) != 15 || parts[10] != "|" {
		return nil, errors.New("invalid entry line format")
	}

	input := parts[:10]
	output := parts[11:]
	entry := Entry{Input: input, Output: output}
	return &entry, nil
}

func ParsePuzzle(lines []string) (Puzzle, error) {
	entries := make([]*Entry, len(lines))
	for i, line := range lines {
		entry, err := ParseEntry(line)
		if err != nil {
			return nil, err
		}

		entries[i] = entry
	}

	return Puzzle(entries), nil
}

func GetDisplay(signal string) Display {
	switch len(signal) {
	case 2:
		return DisplayOne
	case 3:
		return DisplaySeven
	case 4:
		return DisplayFour
	case 7:
		return DisplayEight
	default:
		return DisplayUnknown
	}
}

func (p Puzzle) GetKnownOutputCount() int {
	count := 0
	for _, entry := range p {
		for _, output := range entry.Output {
			if GetDisplay(output) != DisplayUnknown {
				count++
			}
		}
	}

	return count
}
