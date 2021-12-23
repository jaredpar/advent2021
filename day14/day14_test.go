package day14

import (
	_ "embed"
	"testing"

	"advent2021.com/util"
)

//go:embed sample.txt
var sampleText string

//go:embed input.txt
var inputText string

func TestPart1(t *testing.T) {
	run := func(text string, expected int) {
		lines := util.ParseLines(text)
		d, err := ParseData(lines)
		if err != nil {
			t.Fatal(err)
		}

		result := Part1(d)
		if result != expected {
			t.Errorf("expected %d but got %d", expected, result)
		}
	}

	run(sampleText, 1588)
	run(inputText, 2947)
}
