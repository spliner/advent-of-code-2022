package model

type Tree struct {
	height int
	x      int
	y      int
}

func NewTree(height, x, y int) *Tree {
	return &Tree{
		height: height,
		x:      x,
		y:      y,
	}
}

func (t *Tree) Height() int {
	return t.height
}

func (t *Tree) X() int {
	return t.x
}

func (t *Tree) Y() int {
	return t.y
}
