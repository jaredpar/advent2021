package day04

import (
	"testing"
)

func TestCheckWinRow(t *testing.T) {
	b := newBoard(3)
	b.setSlot(0, 0, 1)
	b.setSlot(0, 1, 2)
	b.setSlot(0, 3, 2)
	result := b.checkWin(0, 0, map[int]bool{0: true, 1: true, 2: true})
	if !result {
		t.Error("Not a win")
	}
}

func TestParsePuzzleSample(t *testing.T) {
	puzzle, err := parsePuzzle("sample.txt", 5)
	if err != nil {
		t.Error("cannot parse puzzle")
	}

	if len(puzzle.boards) != 3 {
		t.Error("wrong number of puzzles")
	}

}

func TsetRunPuzzleSample(t *testing.T) {
	result := GetFirstScoreFromFile("sample.txt", 5)
	if result != 4512 {
		t.Error("Bad value detected")
	}

}
