package tree

import (
	"fmt"
	"strconv"

	"github.com/fogleman/gg"
)

type Node struct {
	Element     int
	Left, Right *Node
	gg.Point
}

func (n *Node) Find(value int) *Node {
	if n == nil {
		return nil
	}

	if n.Element < value {
		return n.Right.Find(value)
	} else if n.Element > value {
		return n.Left.Find(value)
	}
	return n
}

func (n *Node) FindMinNode() *Node {
	if n == nil {
		return nil
	}

	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (n *Node) FindMaxNode() *Node {
	if n == nil {
		return nil
	}
	if n.Right == nil {
		return n
	}
	return n.Right.FindMaxNode()
}

func (n *Node) Insert(value int) *Node {
	if n == nil {
		return &Node{Element: value}
	}
	if n.Element > value {
		n.Left = n.Left.Insert(value)
	} else if n.Element < value {
		n.Right = n.Right.Insert(value)
	}
	return n
}

func (n *Node) Delete(value int) bool {
	if n == nil {
		return false
	}

	current := n
	parent := n
	var findValue int
	for current != nil {
		if current.Element < value {
			parent = current
			current = current.Right
		} else if current.Element > value {
			parent = current
			current = current.Left
		} else {
			if current.Left != nil && current.Right != nil {
				temp := current.Right.FindMinNode()
				current.Element = temp.Element
				parent = current
				current = current.Right
				value = temp.Element
			} else {
				findValue = current.Element
				if current.Left == nil && current.Right != nil {
					parent.Left = current.Right
					current = nil
				} else if current.Right == nil && current.Left != nil {
					parent.Right = current.Left
					current = nil
				} else {
					parent.DeleteLeaf(value)
					current = nil
				}
			}
		}
	}

	return findValue == value
}

func (n *Node) DeleteLeaf(value int) {
	if n.Left != nil && n.Left.Element == value {
		n.Left = nil
	} else if n.Right != nil && n.Right.Element == value {
		n.Right = nil
	}
}

func (n *Node) Print() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Element)
}

func (n *Node) RetrieveFunc(f func(n *Node)) {
	if n == nil {
		return
	}
	n.Left.RetrieveFunc(f)
	f(n)
	n.Right.RetrieveFunc(f)
}

func (n *Node) Draw(dc *gg.Context, draw *Draw) {
	dc.SetRGBA(0, 0, 0, 1)
	dc.DrawStringAnchored(strconv.Itoa(n.Element), n.X, n.Y, 0.5, 0.5)
	dc.SetRGBA(0, 0, 0, 1)
	if n.Left != nil {
		n.Left.Point = gg.Point{X: n.X - draw.NodeDistance, Y: n.Y + draw.NodeDistance}
		dc.SetLineWidth(1)
		dc.DrawLine(n.X, n.Y+draw.NodeRadio.X, n.Left.X, n.Left.Y-draw.NodeRadio.Y)
		dc.Stroke()
		n.Left.Draw(dc, draw)
	}
	if n.Right != nil {
		n.Right.Point = gg.Point{X: n.X + draw.NodeDistance, Y: n.Y + draw.NodeDistance}
		dc.SetLineWidth(1)
		dc.DrawLine(n.X, n.Y+draw.NodeRadio.X, n.Right.X, n.Right.Y-draw.NodeRadio.Y)
		dc.Stroke()
		n.Right.Draw(dc, draw)
	}
}
