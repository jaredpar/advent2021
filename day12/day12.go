package day12

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"advent2021.com/util"
)

const labelStart = "start"
const labelEnd = "end"

type Cave struct {
	Label       string
	Connections []*Cave
}

func findCave(label string, caves []*Cave) *Cave {
	for _, cave := range caves {
		if cave.Label == label {
			return cave
		}
	}

	return nil
}

func NewCave(label string) *Cave {
	return &Cave{Label: label}
}

func (c *Cave) IsSmall() bool {
	for _, r := range c.Label {
		if !unicode.IsLower(r) {
			return false
		}
	}

	return true
}

func (c *Cave) IsBig() bool {
	return !c.IsSmall()
}

func (c *Cave) IsStart() bool {
	return c.Label == labelStart
}

func (c *Cave) IsEnd() bool {
	return c.Label == labelEnd
}

func (c *Cave) FindConnection(label string) *Cave {
	return findCave(label, c.Connections)
}

func (c *Cave) AddConnection(other *Cave) {
	if c.FindConnection(other.Label) == nil {
		c.Connections = append(c.Connections, other)
	}

	if other.FindConnection(c.Label) == nil {
		other.Connections = append(other.Connections, c)
	}
}

type CaveSystem struct {
	Caves []*Cave
	Start *Cave
	End   *Cave
}

func NewCaveSystem() *CaveSystem {
	start := NewCave(labelStart)
	end := NewCave(labelEnd)
	caves := []*Cave{start, end}
	return &CaveSystem{Start: start, End: end, Caves: caves}
}

func (cs *CaveSystem) FindCave(label string) *Cave {
	return findCave(label, cs.Caves)
}

func (cs *CaveSystem) AddCave(label string) *Cave {
	c := NewCave(label)
	cs.Caves = append(cs.Caves, c)
	return c
}

func ParseCaveSystem(lines []string) (*CaveSystem, error) {
	cs := NewCaveSystem()

	findOrCreate := func(label string) *Cave {
		c := cs.FindCave(label)
		if c == nil {
			c = cs.AddCave(label)
		}

		return c
	}

	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return nil, errors.New(fmt.Sprintf("string has %d parts: %s", len(parts), line))
		}

		left := findOrCreate(parts[0])
		right := findOrCreate(parts[1])
		left.AddConnection(right)
	}

	return cs, nil
}

func Part1(cs *CaveSystem) []string {
	visited := make(map[string]bool)
	paths := make([]string, 0)
	var visitOne func(c *Cave, path string)
	visitOne = func(c *Cave, path string) {
		if len(path) > 0 {
			path = path + "," + c.Label
		} else {
			util.Assert(c.Label == labelStart)
			path = c.Label
		}

		if c.IsEnd() {
			paths = append(paths, path)
			return
		}

		if c.IsSmall() {
			_, present := visited[c.Label]
			if present {
				return
			}

			visited[c.Label] = true
			defer delete(visited, c.Label)
		}

		for _, connection := range c.Connections {
			if !connection.IsStart() {
				visitOne(connection, path)
			}
		}
	}

	visitOne(cs.Start, "")
	return paths
}

func Part2(cs *CaveSystem) []string {
	visited := make(map[string]bool)
	inDoubleVisit := false

	paths := make([]string, 0)
	var visitOne func(c *Cave, path string)
	visitOne = func(c *Cave, path string) {
		if len(path) > 0 {
			path = path + "," + c.Label
		} else {
			util.Assert(c.Label == labelStart)
			path = c.Label
		}

		if c.IsEnd() {
			paths = append(paths, path)
			return
		}

		if c.IsSmall() {
			_, present := visited[c.Label]
			if present {
				if !inDoubleVisit {
					inDoubleVisit = true
					defer func() {
						inDoubleVisit = false
					}()
				} else {
					return
				}
			} else {
				visited[c.Label] = true
				defer delete(visited, c.Label)
			}
		}

		for _, connection := range c.Connections {
			if !connection.IsStart() {
				visitOne(connection, path)
			}
		}
	}

	visitOne(cs.Start, "")
	return paths
}
