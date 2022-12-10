package model

import "math"

type Rope struct {
	head *Position
	tail *Position
}

func NewRope() *Rope {
	return &Rope{
		head: &Position{
			X: 0,
			Y: 0,
		},
		tail: &Position{
			X: 0,
			Y: 0,
		},
	}
}

func (r *Rope) Head() *Position {
	return r.head
}

func (r *Rope) Tail() *Position {
	return r.tail
}

func (r *Rope) onHeadMove() {
	horizontalDistance := r.head.X - r.tail.X
	verticalDistance := r.head.Y - r.tail.Y

	absHorizontalDistance := math.Abs(float64(horizontalDistance))
	absVerticalDistance := math.Abs(float64(verticalDistance))
	if math.Abs(float64(horizontalDistance)) <= 1 &&
		math.Abs(float64(verticalDistance)) <= 1 {
		return
	}

	var newX, newY int
	if absHorizontalDistance > absVerticalDistance {
		if horizontalDistance > 0 {
			newX = r.head.X - 1
		} else {
			newX = r.head.X + 1
		}
		newY = r.head.Y
	} else if absVerticalDistance > absHorizontalDistance {
		if verticalDistance > 0 {
			newY = r.head.Y - 1
		} else {
			newY = r.head.Y + 1
		}
		newX = r.head.X
	}
	r.tail.X = newX
	r.tail.Y = newY
}

func (r *Rope) MoveRight() {
	r.head.X++
	r.onHeadMove()
}

func (r *Rope) MoveLeft() {
	r.head.X--
	r.onHeadMove()
}

func (r *Rope) MoveUp() {
	r.head.Y++
	r.onHeadMove()
}

func (r *Rope) MoveDown() {
	r.head.Y--
	r.onHeadMove()
}
