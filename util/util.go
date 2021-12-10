package util

import (
	"bufio"
	"embed"
	"strconv"
	"unicode"
)

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

func Assert(cond bool) {
	if !cond {
		panic("failed assert")
	}
}
