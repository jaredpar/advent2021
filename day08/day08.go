package day08

import (
	"errors"
	"fmt"
	"math/bits"
	"sort"
	"strings"

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

const allLetters = "abcdegf"

type letterBits uint

func getBitValue(b byte) letterBits {
	shift := letterBits('g' - b)
	return 1 << shift
}

func convertToBits(s string) letterBits {
	bits := letterBits(0)
	for i := 0; i < len(s); i++ {
		bits |= getBitValue(s[i])
	}

	return bits
}

func getDisplayForBits(l letterBits) int {
	switch l {
	case convertToBits("abcefg"):
		return 0
	case convertToBits("cf"):
		return 1
	case convertToBits("acdeg"):
		return 2
	case convertToBits("acdfg"):
		return 3
	case convertToBits("bcdf"):
		return 4
	case convertToBits("abdfg"):
		return 5
	case convertToBits("abdefg"):
		return 6
	case convertToBits("acf"):
		return 7
	case convertToBits("abcdefg"):
		return 8
	case convertToBits("abcdfg"):
		return 9
	default:
		panic(fmt.Sprintf("bad input: %s", l.String()))
	}
}

func (l letterBits) String() string {
	var sb strings.Builder
	for i := 0; i < len(allLetters); i++ {
		current := allLetters[i]
		if l&getBitValue(current) != 0 {
			sb.WriteByte(current)
		}
	}

	return sb.String()
}

func getSingleLetter(b letterBits) (letter byte, ok bool) {
	letter = 0
	for i := 0; i < len(allLetters); i++ {
		current := allLetters[i]
		if b&getBitValue(current) != 0 {
			if letter != 0 {
				ok = false
				return
			} else {
				letter = current
			}
		}
	}

	ok = true
	return
}

func onesCount(b letterBits) int {
	return bits.OnesCount8(uint8(b))
}

func learnMapping(inputs []string) (map[byte]byte, error) {

	// `possible` is a map of a scrambled letter to possible unscrambled
	// letters
	possible := make(map[byte]letterBits)
	allBits := convertToBits(allLetters)
	for i := 0; i < len(allLetters); i++ {
		possible[allLetters[i]] = allBits
	}

	// `solved` is a map of scrambled letter to unscrambled letters
	solved := make(map[byte]byte)

	// Set of letters solved on the current iteration
	newSolved := make([]byte, 0)

	// Learn that the scrambled `letter` can only map to the specified
	// setters
	learnLetter := func(fromLetter byte, toLetters letterBits) bool {
		old, _ := possible[fromLetter]
		new := old & toLetters
		if old == new {
			return false
		}

		possible[fromLetter] = new
		if onesCount(new) == 1 {
			newSolved = append(newSolved, fromLetter)
		}

		return true
	}

	learnLetterNot := func(fromLetter byte, toLetters letterBits) bool {
		old, _ := possible[fromLetter]
		return learnLetter(fromLetter, old&(^toLetters))
	}

	learnLettersNot := func(fromLetters letterBits, toLetters string) bool {
		progress := false
		toBits := convertToBits(toLetters)
		for i := 0; i < len(allLetters); i++ {
			fromLetter := allLetters[i]
			if (getBitValue(fromLetter) & fromLetters) != 0 {
				progress = learnLetterNot(fromLetter, toBits) || progress
			}
		}

		return progress
	}

	learnLetterSet := func(fromLetters string, toLetters string) bool {
		progress := false
		toBits := convertToBits(toLetters)
		for i := 0; i < len(fromLetters); i++ {
			fromLetter := fromLetters[i]
			progress = learnLetter(fromLetter, toBits) || progress
		}

		return progress
	}

	// Use knowledge about the length of the input to help us reduce the
	// possible values of letters. The `inputs` slice will be ordered before
	// this function is called
	// Len | Digits
	//   2 | 1
	//   3 | 7
	//   4 | 4
	//   5 | 2, 3, 5
	//   6 | 0, 6, 9
	//   7 | 8
	reduceByLength := func() bool {

		progress := false

		for i := 0; i < len(inputs); i++ {
			input := inputs[i]
			switch len(input) {
			case 2:
				progress = learnLetterSet(input, "cf") || progress
			case 3:
				progress = learnLetterSet(input, "acf") || progress
			case 4:
				progress = learnLetterSet(input, "bcdf") || progress
			case 5:
				if i+2 < len(inputs) && len(inputs[i+1]) == 5 && len(inputs[i+2]) == 5 {
					commonBits := convertToBits(input) & convertToBits(inputs[i+1]) & convertToBits(inputs[i+2])
					progress = learnLettersNot(commonBits, "bcef") || progress
					progress = learnLettersNot(^commonBits, "adg") || progress
				}
			case 6:
				if i+2 < len(inputs) && len(inputs[i+1]) == 6 && len(inputs[i+2]) == 6 {
					commonBits := convertToBits(input) & convertToBits(inputs[i+1]) & convertToBits(inputs[i+2])
					progress = learnLettersNot(commonBits, "cde") || progress
					progress = learnLettersNot(^commonBits, "abfg") || progress
				}
			}
		}

		return progress
	}

	reduceSolved := func() bool {
		if len(newSolved) == 0 {
			return false
		}

		temp := newSolved
		newSolved = make([]byte, 0)
		for _, fromLetter := range temp {
			toLetterBits := possible[fromLetter]
			toLetter, ok := getSingleLetter(toLetterBits)
			if !ok {
				panic("logic error")
			}
			solved[fromLetter] = toLetter

			for i := 0; i < len(allLetters); i++ {
				current := allLetters[i]
				if current != fromLetter {
					learnLetterNot(current, toLetterBits)
				}
			}
		}

		return true
	}

	sort.Slice(inputs, func(i, j int) bool { return len(inputs[i]) < len(inputs[j]) })

	for {
		progress := reduceByLength()
		progress = reduceSolved() || progress

		if len(solved) == len(allLetters) {
			// dumpMap(possible)
			break
		}

		if !progress {
			return solved, errors.New("could not resolve")
		}
	}

	return solved, nil
}

func dumpMap(m map[byte]letterBits) {
	for i := 0; i < len(allLetters); i++ {
		current := allLetters[i]
		value := m[current]
		fmt.Printf("%s - %s\n", string(current), value.String())
	}
}

func translateText(text string, letterMap map[byte]byte) letterBits {
	l := letterBits(0)
	for i := 0; i < len(text); i++ {
		current := text[i]
		current, found := letterMap[current]
		if !found {
			panic("map is incomplete")
		}
		l |= getBitValue(current)
	}

	return l
}

func solveEntry(entry *Entry) (int, error) {
	letterMap, err := learnMapping(entry.Input)
	if err != nil {
		return 0, err
	}

	value := 0
	for _, output := range entry.Output {
		value *= 10
		outputBits := translateText(output, letterMap)
		value += getDisplayForBits(outputBits)
	}

	return value, nil
}

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

func GetDisplayEasy(signal string) Display {
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

func (p Puzzle) GetKnownEasyOutputCount() int {
	count := 0
	for _, entry := range p {
		for _, output := range entry.Output {
			if GetDisplayEasy(output) != DisplayUnknown {
				count++
			}
		}
	}

	return count
}

func (p Puzzle) Solve() (int, error) {
	value := 0
	for _, entry := range p {
		v, err := solveEntry(entry)
		if err != nil {
			return 0, err
		}

		value += v
	}

	return value, nil
}
