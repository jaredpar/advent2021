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

	count := d.board.GetOverlapCount()
	if count != 5 {
		t.Error("Invalid count")
	}
}
