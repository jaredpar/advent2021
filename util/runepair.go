package util

import "strings"

type RunePair struct {
	Left, Right rune
}

func NewRunePair(left, right rune) RunePair {
	return RunePair{Left: left, Right: right}
}

func (p RunePair) String() string {
	var sb strings.Builder
	sb.WriteRune(p.Left)
	sb.WriteRune(p.Right)
	return sb.String()
}
