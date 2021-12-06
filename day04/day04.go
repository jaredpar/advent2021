package day04

import (
	"container/list"
	"embed"
	"errors"
	"strconv"
	"strings"

	"advent2021.com/util"
)

//go:embed *.txt
var f embed.FS

type board struct {
	slots    []int
	sideSize int
}

func newBoard(sideSize int) board {
	slots := make([]int, sideSize*sideSize)
	return board{slots: slots, sideSize: sideSize}
}

func (b board) getSlotForValue(value int) (row int, column int) {
	for i, v := range b.slots {
		if v == value {
			row = i / b.sideSize
			column = i % b.sideSize
			return
		}
	}

	row = -1
	column = -1
	return
}

func (b board) getSlot(row int, column int) int {
	index := (row * b.sideSize) + column
	return b.slots[index]
}

func (b board) setSlot(row int, column int, slot int) {
	index := (row * b.sideSize) + column
	b.slots[index] = slot
}

func (b board) isSlotFound(row int, column int, foundMap map[int]bool) bool {
	slot := b.getSlot(row, column)
	_, found := foundMap[slot]
	return found
}

func (b board) checkWin(row int, column int, foundMap map[int]bool) bool {
	checkRow := func() bool {
		for c := 0; c < b.sideSize; c++ {
			if !b.isSlotFound(row, c, foundMap) {
				return false
			}
		}

		return true
	}

	checkColumn := func() bool {
		for r := 0; r < b.sideSize; r++ {
			if !b.isSlotFound(r, column, foundMap) {
				return false
			}
		}

		return true
	}

	/*
		checkDiag := func() bool {
			right := 0
			left := 0
			for i := 0; i < b.sideSize; i++ {
				j := b.sideSize - i
				if b.isSlotFound(i, i, foundMap) {
					right++
				}
				if b.isSlotFound(j, j, foundMap) {
					left++
				}
			}

			return left == b.sideSize || right == b.sideSize
		}
	*/

	return checkColumn() || checkRow()
}

func parseBoard(lines []string) (board, error) {
	sideSize := len(lines)
	b := newBoard(sideSize)

	for row, line := range lines {
		values := util.SplitOnWhiteSpace(line)
		if len(values) != sideSize {
			return b, errors.New("incorrect number of columns")
		}

		for column, value := range values {
			v, err := strconv.Atoi(value)
			if err != nil {
				return b, err
			}

			b.setSlot(row, column, v)
		}
	}

	return b, nil
}

type puzzle struct {
	numbers []int
	boards  []board
}

func newPuzzle(numbers []int, boards []board) *puzzle {
	return &puzzle{numbers: numbers, boards: boards}
}

func parsePuzzle(fileName string, sideSize int) (*puzzle, error) {
	parseNumbers := func(line string) ([]int, error) {
		items := strings.Split(line, ",")
		numbers := make([]int, len(items))

		for i, item := range items {
			number, err := strconv.Atoi(item)
			if err != nil {
				return nil, err
			}

			numbers[i] = number
		}

		return numbers, nil
	}

	parseBoards := func(lines []string) ([]board, error) {
		index := 0
		boards := make([]board, 0)
		for {
			if index >= len(lines) {
				return boards, nil
			}

			b, err := parseBoard(lines[index : index+sideSize])
			if err != nil {
				return nil, err
			}

			boards = append(boards, b)
			index += sideSize + 1
		}
	}

	lines, err := util.ReadLines(f, fileName)
	if err != nil {
		return nil, err
	}

	if len(lines) < 2+sideSize {
		return nil, err
	}

	numbers, err := parseNumbers(lines[0])
	if err != nil {
		return nil, err
	}

	boards, err := parseBoards(lines[2:])
	if err != nil {
		return nil, err
	}

	return newPuzzle(numbers, boards), nil
}

func runPuzzle(p *puzzle) int {

	m := make(map[int]bool)
	for _, current := range p.numbers {
		m[current] = true
		for _, b := range p.boards {
			row, column := b.getSlotForValue(current)
			if row >= 0 && b.checkWin(row, column, m) {
				sum := 0
				for _, value := range b.slots {
					if _, found := m[value]; !found {
						sum += value
					}
				}

				return sum * current
			}
		}
	}

	return -1
}

func RunPuzzleFile(fileName string, sideSize int) int {
	p, err := parsePuzzle(fileName, sideSize)
	if err != nil {
		panic(err)
	}

	return runPuzzle(p)
}

func runPuzzleLastWin(p *puzzle) int {

	m := make(map[int]bool)
	l := list.New()
	for _, b := range p.boards {
		l.PushBack(b)
	}

	for _, current := range p.numbers {
		m[current] = true

		for n := l.Front(); n != nil; {
			b := n.Value.(board)
			row, column := b.getSlotForValue(current)
			if row >= 0 && b.checkWin(row, column, m) {
				temp := n
				n = n.Next()
				l.Remove(temp)

				if l.Len() == 1 {
					sum := 0
					for _, value := range b.slots {
						if _, found := m[value]; !found {
							sum += value
						}
					}

					return sum * current
				}
			} else {
				n = n.Next()
			}
		}
	}

	return -1
}

func RunPuzzleFileLastWin(fileName string, sideSize int) int {
	p, err := parsePuzzle(fileName, sideSize)
	if err != nil {
		panic(err)
	}

	return runPuzzleLastWin(p)
}
