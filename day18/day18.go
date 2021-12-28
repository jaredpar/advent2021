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
	n.Value = 0
	n.Left = nil
	n.Right = nil

	// Fix up the left node
	leftSib := func() *Node {
		cur := n
		for cur != nil {
			if cur.Parent != nil && cur.Parent.Left != cur {
				cur = cur.Parent.Left
				for !cur.IsLeaf() {
					cur = cur.Right
				}
				return cur
			}

			cur = cur.Parent
		}

		return nil
	}()

	if leftSib != nil {
		leftSib.Value += left
	}

	rightSib := func() *Node {
		cur := n
		for cur != nil {
			if cur.Parent != nil && cur.Parent.Right != cur {
				cur = cur.Parent.Right
				for !cur.IsLeaf() {
					cur = cur.Left
				}
				return cur
			}

			cur = cur.Parent
		}

		return nil
	}()

	if rightSib != nil {
		rightSib.Value += right
	}

	util.Assert(n.IsLeaf())
}

func (n *Node) Split() {
	util.Assert(n.IsLeaf())

	extra := n.Value % 2
	n.Left = NewNode((n.Value-extra)/2, n)
	n.Right = NewNode(((n.Value-extra)/2)+extra, n)
	n.Value = -1
}

func (n *Node) Reduce() {

	reduceOne := func() bool {
		explode := n.Find(func(n *Node) bool {
			return n.IsPair() && n.Depth() >= 4
		})

		if explode != nil {
			explode.Explode()
			return true
		}

		split := n.Find(func(n *Node) bool {
			return n.Value > 9
		})

		if split != nil {
			split.Split()
			return true
		}

		return false
	}

	for reduceOne() {
		// fmt.Println(n.String())
	}
}

func (n *Node) Join(other *Node) *Node {
	p := NewNode(-1, nil)
	p.Left = n
	n.Parent = p
	p.Right = other
	other.Parent = p

	p.Reduce()
	return p
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

func (n *Node) Depth() int {
	depth := 0
	cur := n
	for !cur.IsRoot() {
		depth++
		cur = cur.Parent
	}

	return depth
}

func (n *Node) Find(predicate func(*Node) bool) *Node {
	toVisit := []*Node{n}
	for len(toVisit) > 0 {
		next := toVisit[0]
		toVisit = toVisit[1:]

		if predicate(next) {
			return next
		}

		if next.Left != nil {
			toVisit = append(toVisit, next.Left)
		}

		if next.Right != nil {
			toVisit = append(toVisit, next.Right)
		}
	}

	return nil
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

func MustParseNode(text string) *Node {
	node, err := ParseNode(text)
	if err != nil {
		panic(err)
	}

	return node
}

func magnitude(n *Node) int {
	if n.IsLeaf() {
		return n.Value
	} else {
		left := 3 * magnitude(n.Left)
		right := 2 * magnitude(n.Right)
		return left + right
	}
}

func Part1(lines []string) int {
	var node *Node
	for _, line := range lines {
		cur := MustParseNode(line)
		if node == nil {
			node = cur
		} else {
			node = node.Join(cur)
		}

		fmt.Println(node.String())
	}

	return magnitude(node)
}
