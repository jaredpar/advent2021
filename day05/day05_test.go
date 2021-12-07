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
	if count != 12 {
		t.Errorf("expected 12 got %d", count)
	}
}

func TestOverlapCountOnInput(t *testing.T) {
	d, err := readDiagram("input.txt")
	if err != nil {
		t.Error("Cannot parse")
		return
	}

	count := d.board.getOverlapCount()
	expected := 19258
	if count != expected {
		t.Errorf("expected %d but got %d", expected, count)
	}
}
