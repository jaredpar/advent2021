package day05

import (
	"embed"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"advent2021.com/util"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func (l line) isDiagonal() bool {
	rise := math.Abs(float64(l.end.y - l.start.y))
	run := math.Abs(float64(l.end.x - l.start.x))
	return 1 == int(rise/run)
}

func newLine(x1 int, y1 int, x2 int, y2 int) *line {
	return &line{
		start: point{x1, y1},
		end:   point{x2, y2}}
}

func parseLines(textLines []string) ([]*line, error) {
	parsePoint := func(text string) (x int, y int, err error) {
		parts := strings.Split(text, ",")
		if len(parts) != 2 {
			err = errors.New("invalid point format")
			return
		}

		x, err = strconv.Atoi(parts[0])
		if err != nil {
			return
		}

		y, err = strconv.Atoi(parts[1])
		return
	}

	lines := make([]*line, 0, len(textLines))
	for _, textLine := range textLines {
		parts := util.SplitOnWhiteSpace(textLine)
		if len(parts) != 3 {
			return nil, errors.New("invalid line format")
		}

		x1, y1, err := parsePoint(parts[0])
		if err != nil {
			return lines, err
		}

		x2, y2, err := parsePoint(parts[2])
		if err != nil {
			return lines, err
		}

		line := newLine(x1, y1, x2, y2)
		lines = append(lines, line)
	}

	return lines, nil
}

type board struct {
	values []int
	size   int
}

// Get the index into a board for a given x and y position
func (b board) getValueIndex(x int, y int) int {
	index := y * b.size
	return index + x
}

func (b board) getValue(x int, y int) int {
	index := b.getValueIndex(x, y)
	return b.values[index]
}

func newBoard(lines []*line) board {
	size := 0
	for _, line := range lines {
		size = util.Max(size, line.start.x)
		size = util.Max(size, line.start.y)
		size = util.Max(size, line.end.x)
		size = util.Max(size, line.end.y)
	}

	size += 1
	b := board{values: make([]int, size*size), size: size}

	for _, line := range lines {
		if line.start.x == line.end.x {
			x := line.start.x
			yStart := util.Min(line.start.y, line.end.y)
			yEnd := util.Max(line.start.y, line.end.y)
			for y := yStart; y <= yEnd; y++ {
				index := b.getValueIndex(x, y)
				b.values[index]++
			}

		} else if line.start.y == line.end.y {
			y := line.start.y
			xStart := util.Min(line.start.x, line.end.x)
			xEnd := util.Max(line.start.x, line.end.x)
			for x := xStart; x <= xEnd; x++ {
				index := b.getValueIndex(x, y)
				b.values[index]++
			}
		} else if line.isDiagonal() {
			xInc := 1
			if line.start.x > line.end.x {
				xInc = -1
			}

			yInc := 1
			if line.start.y > line.end.y {
				yInc = -1
			}

			p := line.start
			for {
				index := b.getValueIndex(p.x, p.y)
				b.values[index]++

				if p.x == line.end.x && p.y == line.end.y {
					break
				}

				p.x += xInc
				p.y += yInc

			}
		}
	}

	return b
}

func (b board) String() string {
	var sb strings.Builder
	for y := 0; y < b.size; y++ {
		for x := 0; x < b.size; x++ {
			value := b.getValue(x, y)
			fmt.Fprintf(&sb, "%d", value)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (b board) getOverlapCount() int {
	count := 0
	for _, value := range b.values {
		if value > 1 {
			count++
		}
	}
	return count
}

type diagram struct {
	lines []*line
	board board
}

func newDiagram(lines []*line) *diagram {
	return &diagram{lines: lines, board: newBoard(lines)}
}

func parseDiagram(textLines []string) (*diagram, error) {
	lines, err := parseLines(textLines)
	if err != nil {
		return nil, err
	}

	return newDiagram(lines), nil
}

//go:embed *.txt
var f embed.FS

func readDiagram(fileName string) (*diagram, error) {
	lines, err := util.ReadLines(f, fileName)
	if err != nil {
		panic(err)
	}

	return parseDiagram(lines)
}

func GetDiagram(fileName string) string {
	d, err := readDiagram(fileName)
	if err != nil {
		panic(err)
	}

	return d.board.String()
}

func GetOverlapCount() int {
	d, err := readDiagram("input.txt")
	if err != nil {
		panic(err)
	}

	return d.board.getOverlapCount()
}
