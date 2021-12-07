package day05

import (
	"testing"
)

func TestOverlapCountOnSample(t *testing.T) {
	d, err := readDiagram("sample.txt")
	if err != nil {
		t.Error("Cannot parse")
		return
	}

	count := d.board.getOverlapCount()
	if count != 5 {
		t.Error("Invalid count")
	}
}

func TestOverlapCountOnInput(t *testing.T) {
	d, err := readDiagram("input.txt")
	if err != nil {
		t.Error("Cannot parse")
		return
	}

	count := d.board.getOverlapCount()
	if count != 6841 {
		t.Error("Invalid count")
	}
}
