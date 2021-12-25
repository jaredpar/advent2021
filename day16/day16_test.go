package day16

import (
	_ "embed"
	"testing"

	"advent2021.com/testUtil"
)

//go:embed sample.txt
var sampleText string

//go:embed input.txt
var inputText string

func TestSimpleParse(t *testing.T) {
	packet, _ := ParsePacket("D2FE28")
	testUtil.AssertEqualInt(t, 6, packet.Version)
	testUtil.AssertEqualInt(t, 4, packet.TypeId)
}
