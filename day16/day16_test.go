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
	packet, err := ParsePacket("D2FE28")
	if err != nil {
		t.Fatal(err)
	}

	literal, ok := packet.(*LiteralPacket)
	if !ok {
		t.Fatal("wrong type")
	}

	testUtil.AssertEqualInt(t, 4, literal.typeId)
	testUtil.AssertEqualInt(t, 2021, literal.Value)
}
