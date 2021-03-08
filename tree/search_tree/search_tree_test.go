package search_tree

import (
	"fmt"
	"testing"

	"github.com/fogleman/gg"

	"dataStructure-go/tree"
)

const (
	Width                = 512
	Height               = 512
	NodeDistance float64 = 50
)

var (
	RootPoint = gg.Point{X: Width / 2, Y: 10}
	NodeRadio = gg.Point{X: 10, Y: 5}
)

func TestSearchTree(t *testing.T) {
	searchTree := NewSearchTree()
	searchTree.Insert(7)
	searchTree.Insert(8)
	searchTree.Insert(2)
	searchTree.Insert(1)
	searchTree.Insert(6)
	searchTree.Insert(5)
	searchTree.Insert(4)
	searchTree.Insert(3)
	searchTree.Retrieve()
	fmt.Println()
	fmt.Printf("max:%+v\n", searchTree.FindMaxNode())
	fmt.Printf("min:%+v\n", searchTree.FindMinNode())
	fmt.Printf("find:5 ==>%+v\n", searchTree.Find(5))
	fmt.Printf("delete:3 ==>%t\n", searchTree.Delete(2))
	searchTree.Retrieve()

	if err := searchTree.Draw(tree.NewDraw(
		Width, Height, NodeDistance, NodeRadio, RootPoint, "binary_tree.png")); err != nil {
		t.Errorf(err.Error())
	}
}
