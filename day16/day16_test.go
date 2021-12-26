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

func TestSimpleLiteralParse(t *testing.T) {
	packet, err := ParsePacket("D2FE28")
	if err != nil {
		t.Fatal(err)
	}

	literal, ok := packet.(*LiteralPacket)
	if !ok {
		t.Fatal("wrong type")
	}

	testUtil.AssertEqualInt(t, 2021, literal.Value)
}

func TestParseBinaryInt(t *testing.T) {

	testCore := func(text string, expected int) {
		value, err := parseBinaryInt([]rune(text))
		if err != nil {
			t.Fatal(err)
		}
		testUtil.AssertEqualInt(t, expected, value)
	}

	testCore("000000000011011", 27)
	testCore("11", 3)
}

func TestSimpleOperatorParse1(t *testing.T) {
	packet, err := ParsePacket("38006F45291200")
	if err != nil {
		t.Fatal(err)
	}

	operator, ok := packet.(*OperatorPacket)
	if !ok {
		t.Fatal("wrong type")
	}

	if 2 != len(operator.Children) {
		t.Fatalf("wrong number of children: %d", len(operator.Children))
	}

	testVal := func(i, expected int) {
		literal, ok := operator.Children[i].(*LiteralPacket)
		if !ok {
			t.Fatal("wrong type")
		}

		testUtil.AssertEqualInt(t, expected, literal.Value)
	}

	testVal(0, 10)
	testVal(1, 20)

}

func TestSimpleOperatorParse2(t *testing.T) {
	packet, err := ParsePacket("EE00D40C823060")
	if err != nil {
		t.Fatal(err)
	}

	operator, ok := packet.(*OperatorPacket)
	if !ok {
		t.Fatal("wrong type")
	}

	if 3 != len(operator.Children) {
		t.Fatalf("wrong number of children: %d", len(operator.Children))
	}

	testVal := func(i, expected int) {
		literal, ok := operator.Children[i].(*LiteralPacket)
		if !ok {
			t.Fatal("wrong type")
		}

		testUtil.AssertEqualInt(t, expected, literal.Value)
	}

	testVal(0, 1)
	testVal(1, 2)
	testVal(2, 3)
}
