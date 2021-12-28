package day18

import (
	"testing"

	"advent2021.com/testUtil"
)

func TestParseNode(t *testing.T) {
	var testCore func(text string)
	testCore = func(text string) {
		node, err := ParseNode(text)
		if err != nil {
			t.Fatal(err)
		}

		var str = node.String()
		testUtil.AssertEqualString(t, text, str)
	}

	testCore("[1,0]")
	testCore("[[1,0],0]")
	testCore("[[[[[9,8],1],2],3],4]")
}
