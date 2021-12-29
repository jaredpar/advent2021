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

// Get the leaf nodes of the tree in BFS order
func (n *Node) GetLeafNodes() []*Node {
	nodes := make([]*Node, 0)
	_ = n.FindDfs(func(n *Node) bool {
		if n.IsLeaf() {
			nodes = append(nodes, n)
		}

		return false
	})
	return nodes
}

func (n *Node) GetRoot() *Node {
	cur := n
	for !cur.IsRoot() {
		cur = cur.Parent
	}

	return cur
}

func (n *Node) Explode() {
	util.Assert(n.IsPair())

	left := n.Left.Value
	right := n.Right.Value
	n.Value = 0
	n.Left = nil
	n.Right = nil

	leafs := n.GetRoot().GetLeafNodes()
	leafIndex := func() int {
		for i, cur := range leafs {
			if cur == n {
				return i
			}
		}

		panic("node not found")
	}()

	if leafIndex > 0 {
		leafs[leafIndex-1].Value += left
	}

	if leafIndex+1 < len(leafs) {
		leafs[leafIndex+1].Value += right
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
		explode := n.FindDfs(func(n *Node) bool {
			return n.IsPair() && n.Depth() >= 4
		})

		if explode != nil {
			explode.Explode()
			return true
		}

		split := n.FindDfs(func(n *Node) bool {
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
			fmt.Fprintf(&sb, "%d", cur.Value)
		} else if cur.IsPair() {
			fmt.Fprintf(&sb, "[%d,%d]", cur.Left.Value, cur.Right.Value)
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

// Run the predicate on the nodes in the tree in BFS order. Will return
// first Node for which the predicate returns true, nil otherwise
func (n *Node) FindBfs(predicate func(*Node) bool) *Node {
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

// Run the predicate on the nodes in the tree in BFS order. Will return
// first Node for which the predicate returns true, nil otherwise
func (n *Node) FindDfs(predicate func(*Node) bool) *Node {
	toVisit := []*Node{n}
	for len(toVisit) > 0 {
		index := len(toVisit) - 1
		cur := toVisit[index]
		toVisit = toVisit[:index]

		if predicate(cur) {
			return cur
		}

		if cur.Right != nil {
			toVisit = append(toVisit, cur.Right)
		}

		if cur.Left != nil {
			toVisit = append(toVisit, cur.Left)
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
