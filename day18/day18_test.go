package day18

import (
	"fmt"
	"strings"
	"testing"

	"advent2021.com/testUtil"
	"advent2021.com/util"
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

func TestFindDfs(t *testing.T) {
	testCore := func(text string, expected string) {
		node := MustParseNode(text)
		var sb strings.Builder
		_ = node.FindDfs(func(n *Node) bool {
			if n.IsLeaf() {
				if sb.Len() > 0 {
					sb.WriteRune(' ')
				}
				fmt.Fprintf(&sb, "%d", n.Value)
			}
			return false
		})

		testUtil.AssertEqualString(t, expected, sb.String())
	}

	testCore("[[[[[9,8],1],2],3],4]", "9 8 1 2 3 4")
}

func TestExplode(t *testing.T) {
	node := MustParseNode("[[[[[9,8],1],2],3],4]")
	target := node.FindBfs(func(n *Node) bool {
		return n.Value == 9
	})

	target.Parent.Explode()
	testUtil.AssertEqualString(t, "[[[[0,9],2],3],4]", node.String())

	node = MustParseNode("[7,[6,[5,[4,[3,2]]]]]")
	target = node.FindBfs(func(n *Node) bool {
		return n.Value == 3
	})

	target.Parent.Explode()
	testUtil.AssertEqualString(t, "[7,[6,[5,[7,0]]]]", node.String())
}

func TestSplit(t *testing.T) {
	node := MustParseNode("7")
	node.Split()
	testUtil.AssertEqualString(t, "[3,4]", node.String())

	node = MustParseNode("6")
	node.Split()
	testUtil.AssertEqualString(t, "[3,3]", node.String())
}

func TestReduce(t *testing.T) {
	testCore := func(original, expected string) {
		node := MustParseNode(original)
		node.Reduce()
		testUtil.AssertEqualString(t, expected, node.String())
	}

	testCore("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
}

func TestMagnitude(t *testing.T) {
	testCore := func(text string, expected int) {
		node := MustParseNode(text)
		node.Reduce()
		actual := magnitude(node)
		testUtil.AssertEqualInt(t, expected, actual)
	}

	testCore("[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137)
	testCore("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488)
	testCore("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488)
}

func TestJoinLines(t *testing.T) {
	const data = `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`

	lines := util.ParseLines(data)
	node := joinLines(lines)
	testUtil.AssertEqualString(t, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", node.String())
}
