package search_tree

import (
	"fmt"

	"github.com/fogleman/gg"

	"dataStructure-go/tree"
)

type SearchTree struct {
	root *tree.Node
}

func NewSearchTree() *SearchTree {
	return &SearchTree{}
}

func (s *SearchTree) Find(value int) *tree.Node {
	return s.root.Find(value)
}

func (s *SearchTree) FindMinNode() *tree.Node {
	return s.root.FindMinNode()
}

func (s *SearchTree) FindMaxNode() *tree.Node {
	return s.root.FindMaxNode()
}

func (s *SearchTree) IsEmpty() bool {
	return s.root == nil
}

func (s *SearchTree) newNode(value int) *tree.Node {
	return &tree.Node{Element: value}
}

func (s *SearchTree) Insert(value int) {
	if s.root == nil {
		s.root = s.newNode(value)
		return
	}
	s.root.Insert(value)
}

func (s *SearchTree) Delete(value int) bool {
	return s.root.Delete(value)
}

func (s *SearchTree) Retrieve() {
	if s.root == nil {
		return
	}
	s.root.RetrieveFunc(func(n *tree.Node) {
		n.Print()
	})
	fmt.Println()
}

func (s *SearchTree) Draw(draw *tree.Draw) error {
	if s.root == nil {
		return nil
	}
	s.root.Point = draw.RootPoint
	dc := gg.NewContext(draw.Width, draw.Height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	s.root.Draw(dc, draw)
	s.root.Left.Draw(dc, draw)
	s.root.Right.Draw(dc, draw)

	if err := dc.SavePNG(draw.Path); err != nil {
		return err
	}
	return nil
}
