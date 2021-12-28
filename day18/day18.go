package day18

import (
	"fmt"
	"strings"

	"advent2021.com/util"
)

type Node struct {
	Value               int
	Parent, Left, Right *Node
}

func NewNode(value int, parent *Node) *Node {
	return &Node{Value: value, Parent: parent, Left: nil, Right: nil}
}

func (n *Node) IsRoot() bool {
	return n.Parent == nil
}

func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node) IsPair() bool {
	return !n.IsLeaf() && n.Left.IsLeaf() && n.Right.IsLeaf()
}

func (n *Node) Explode() {
	util.Assert(n.IsPair())

	left := n.Left.Value
	right := n.Right.Value
	n.Left = nil
	n.Right = nil

	// Fix up the left node
	if n.Parent.Left != nil && n.Parent.Left != n {
		cur := n.Parent.Left
		for !cur.IsPair() {
			cur = cur.Right
		}

		cur.Right.Value += left
	}

	// Fix up the right node
	if n.Parent.Right != nil && n.Parent.Right != n {
		cur := n.Parent.Right
		for !cur.IsPair() {
			cur = cur.Left
		}

		cur.Left.Value += right
	}

	util.Assert(n.IsLeaf())
}

func (n *Node) String() string {
	var sb strings.Builder
	var impl func(*Node)
	impl = func(cur *Node) {
		if cur.IsLeaf() {
			sb.WriteString(fmt.Sprintf("%d", cur.Value))
		} else if cur.IsPair() {
			sb.WriteString(fmt.Sprintf("[%d,%d]", cur.Left.Value, cur.Right.Value))
		} else {
			sb.WriteString("[")
			impl(cur.Left)
			sb.WriteString(",")
			impl(cur.Right)
			sb.WriteString("]")
		}
	}
	impl(n)
	return sb.String()
}

func ParseNode(text string) (*Node, error) {
	var impl func(*Node) (*Node, error)
	index := 0
	impl = func(parent *Node) (*Node, error) {
		if index == len(text) {
			return nil, nil
		}

		var err error
		if text[index] == '[' {
			node := NewNode(-1, parent)
			index++

			node.Left, err = impl(node)
			if err != nil {
				return nil, err
			}

			if text[index] != ',' {
				return nil, fmt.Errorf("expected ']' got '%b'", text[index])
			}
			index++

			node.Right, err = impl(node)
			if err != nil {
				return nil, err
			}

			if text[index] != ']' {
				return nil, fmt.Errorf("expected ']' got '%b'", text[index])
			}
			index++

			return node, nil
		} else {
			var value int
			value, err = util.ByteToInt(text[index])
			if err != nil {
				return nil, err
			}
			index++

			return NewNode(value, parent), nil
		}
	}

	return impl(nil)
}
