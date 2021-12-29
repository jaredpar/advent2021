package util

import (
	"bufio"
	"embed"
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func ParseLines(text string) []string {
	reader := strings.NewReader(text)
	scanner := bufio.NewScanner(reader)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadLines(f embed.FS, name string) ([]string, error) {
	file, err := f.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func MustReadLines(f embed.FS, name string) []string {
	lines, err := ReadLines(f, name)
	if err != nil {
		panic(err)
	}

	return lines
}

func ReadLinesAsInt(f embed.FS, name string) ([]int, error) {
	lines, err := ReadLines(f, name)
	if err != nil {
		return nil, err
	}

	values := make([]int, 0, len(lines))
	for _, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	return values, nil
}

func ReadAsSingleLine(f embed.FS, name string) (string, error) {
	file, err := f.Open(name)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return "", nil
	}

	line := scanner.Text()
	if scanner.Scan() {
		return "", errors.New("file had multiple lines")
	}

	return line, nil
}

func ParseCommaSepInt(line string) ([]int, error) {
	parts := strings.Split(line, ",")
	numbers := make([]int, 0, len(parts))
	for _, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func SplitOnWhiteSpace(line string) []string {
	startIndex := -1
	items := make([]string, 0)
	for index, r := range line {
		if unicode.IsSpace(r) {
			if startIndex >= 0 {
				item := line[startIndex:index]
				items = append(items, item)
			}
			startIndex = -1
		} else if startIndex < 0 {
			startIndex = index
		}
	}

	if startIndex >= 0 {
		item := line[startIndex:]
		items = append(items, item)
	}

	return items
}

// Convert a rune value between '0' and '9' to a int value
func RuneToInt(r rune) (int, error) {
	value := int(r - '0')
	if value >= 0 && value <= 9 {
		return value, nil
	}

	return value, errors.New("invalid value")
}

// Convert a rune value between '0' and '9' to a int value
func ByteToInt(b byte) (int, error) {
	value := int(b - '0')
	if value >= 0 && value <= 9 {
		return value, nil
	}

	return value, errors.New("invalid value")
}

// Return the first rune in the string. Will panic on a zero length string
func FirstRune(text string) rune {
	for _, r := range text {
		return r
	}

	panic("zero length string")
}

func SetAllInt(values []int, value int) {
	for i := 0; i < len(values); i++ {
		values[i] = value
	}
}

func DigitToRune(d int) rune {
	return rune('0' + d)
}

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		x *= -1
	}

	return x
}

func Require(cond bool) {
	if !cond {
		panic("failed assert")
	}
}
