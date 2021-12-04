package day03

import (
	"container/list"
	"embed"
	"errors"
	"math"
	"strconv"

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

func GetLifeSupportRates(fileName string) (oxygen int, co2 int, err error) {
	lines, err := util.ReadLines(f, fileName)
	if err != nil {
		return 0, 0, err
	}

	if len(lines) == 0 {
		return 0, 0, nil
	}

	search := func(useMostCommon bool) (int, error) {

		l := list.New()
		for _, line := range lines {
			l.PushBack(line)
		}

		getSearchValue := func(column int) byte {
			counter := 0
			for current := l.Front(); current != nil; current = current.Next() {
				line := current.Value.(string)
				if line[column] == '1' {
					counter++
				} else {
					counter--
				}
			}

			if useMostCommon {
				if counter >= 0 {
					return '1'
				} else {
					return '0'
				}
			} else {
				if counter >= 0 {
					return '0'
				} else {
					return '1'
				}
			}
		}

		filterCount := 0
		for column := 0; column < len(lines[0]); column++ {
			searchValue := getSearchValue(column)
			current := l.Front()

			for {
				if current == nil {
					break
				}

				line := current.Value.(string)
				if line[column] != searchValue {
					next := current.Next()
					l.Remove(current)
					current = next
					filterCount++
				} else {
					current = current.Next()
				}
			}

			if filterCount+1 == len(lines) {
				line := l.Front().Value.(string)
				longValue, err := strconv.ParseInt(line, 2, 32)
				if err != nil {
					return 0, err
				}

				return int(longValue), nil
			}
		}

		return 0, errors.New("search failed")
	}

	oxygen, err = search(true)
	if err != nil {
		return 0, 0, err
	}

	co2, err = search(false)
	return
}

func GetLifeSupportRatesValue(fileName string) int {
	oxygen, co2, err := GetLifeSupportRates(fileName)
	if err != nil {
		panic(err)
	}

	return oxygen * co2
}
