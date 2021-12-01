package util

import (
	"bufio"
	"embed"
	"strconv"
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
