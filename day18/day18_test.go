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

func TestExplode(t *testing.T) {
	node := MustParseNode("[[[[[9,8],1],2],3],4]")
	target := node.Find(func(n *Node) bool {
		return n.Value == 9
	})

	target.Parent.Explode()
	testUtil.AssertEqualString(t, "[[[[0,9],2],3],4]", node.String())

	node = MustParseNode("[7,[6,[5,[4,[3,2]]]]]")
	target = node.Find(func(n *Node) bool {
		return n.Value == 3
	})

	target.Parent.Explode()
	testUtil.AssertEqualString(t, "[7,[6,[5,[7,0]]]]", node.String())
}