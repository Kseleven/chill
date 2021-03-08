package main

import (
	"dataStructure-go/tree"
	"strconv"

	"github.com/fogleman/gg"
)

const (
	Width                = 512
	Height               = 512
	NodeDistance float64 = 100
)

var (
	RootPoint = gg.Point{X: Width / 2, Y: 10}
)

type NodeDraw struct {
	*tree.Node
}

func (n *NodeDraw) RetrieveDraw() {
	if n == nil {
		return
	}
	n.Point = RootPoint
	dc := gg.NewContext(Width, Height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	n.Draw(dc, NodeDistance)
	n.Left.Draw(dc, NodeDistance)
	n.Right.Draw(dc, NodeDistance)

	if err := dc.SavePNG("image.png"); err != nil {
		panic(err)
	}
}

func main() {
	dc := gg.NewContext(Width, Height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	beginPoint := gg.Point{X: Width / 2, Y: 10}
	nextPoint := beginPoint
	var distance float64 = 100
	for i := 0; i < 2; i++ {
		nextPoint = gg.Point{X: beginPoint.X + distance, Y: beginPoint.Y + distance}
		dc.SetRGBA(0, 0, 0, 1)
		dc.DrawStringAnchored(strconv.Itoa(i+1), beginPoint.X, beginPoint.Y, 0.5, 0.5)
		dc.SetRGBA(0, 0, 0, 1)
		dc.SetLineWidth(1)
		dc.DrawLine(beginPoint.X, beginPoint.Y+5, nextPoint.X, nextPoint.Y)
		dc.Stroke()
		beginPoint = nextPoint
	}

	if err := dc.SavePNG("image.png"); err != nil {
		panic(err)
	}
}
