package tree

import "github.com/fogleman/gg"

type Draw struct {
	Width, Height int
	NodeDistance  float64
	NodeRadio     gg.Point
	RootPoint     gg.Point
	Path          string
}

func NewDraw(width, height int, nodeDistance float64, nodeRadio, rootPoint gg.Point, path string) *Draw {
	return &Draw{
		Width:        width,
		Height:       height,
		NodeDistance: nodeDistance,
		NodeRadio:    nodeRadio,
		RootPoint:    rootPoint,
		Path:         path,
	}
}
