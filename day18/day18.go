package day18

import (
	"math"

	"advent2021.com/util"
)

func magnitude(n *Node) int {
	if n.IsLeaf() {
		return n.Value
	} else {
		left := 3 * magnitude(n.Left)
		right := 2 * magnitude(n.Right)
		return left + right
	}
}

func joinLines(lines []string) *Node {
	var node *Node
	for _, line := range lines {
		cur := MustParseNode(line)
		if node == nil {
			node = cur
		} else {
			node = node.Join(cur)
		}

		// fmt.Println(node.String())
	}

	return node
}

func Part1(lines []string) int {
	node := joinLines(lines)
	return magnitude(node)
}

func Part2(lines []string) int {
	max := math.MinInt
	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines); y++ {
			left := MustParseNode(lines[x])
			right := MustParseNode(lines[y])
			magnitude := magnitude(left.Join(right))
			max = util.Max(max, magnitude)
		}
	}
	return max
}
