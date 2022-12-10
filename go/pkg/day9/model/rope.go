package model

import (
	"math"
)

type Rope struct {
	knots []*Position
}

func NewRope(knotCount int) *Rope {
	knots := make([]*Position, knotCount)
	for i := 0; i < knotCount; i++ {
		knots[i] = &Position{0, 0}
	}
	return &Rope{
		knots: knots,
	}
}

func (r *Rope) Head() *Position {
	return r.knots[0]
}

func (r *Rope) Tail() *Position {
	return r.knots[len(r.knots)-1]
}

func (r *Rope) onHeadMove() {
	previousKnot := r.Head()
	for _, knot := range r.knots[1:] {
		onMove(previousKnot, knot)
		previousKnot = knot
	}
}

func onMove(previousKnot, knot *Position) {
	horizontalDistance := previousKnot.X - knot.X
	verticalDistance := previousKnot.Y - knot.Y

	absHorizontalDistance := math.Abs(float64(horizontalDistance))
	absVerticalDistance := math.Abs(float64(verticalDistance))
	if math.Abs(float64(horizontalDistance)) <= 1 &&
		math.Abs(float64(verticalDistance)) <= 1 {
		return
	}

	var newX, newY int
	if absVerticalDistance > absHorizontalDistance {
		if verticalDistance > 0 {
			newY = previousKnot.Y - 1
		} else {
			newY = previousKnot.Y + 1
		}
		newX = previousKnot.X
	} else if absHorizontalDistance > absVerticalDistance {
		if horizontalDistance > 0 {
			newX = previousKnot.X - 1
		} else {
			newX = previousKnot.X + 1
		}
		newY = previousKnot.Y
	} else {
		if verticalDistance > 0 {
			newY = previousKnot.Y - 1
		} else {
			newY = previousKnot.Y + 1
		}
		if horizontalDistance > 0 {
			newX = previousKnot.X - 1
		} else {
			newX = previousKnot.X + 1
		}
	}
	knot.X = newX
	knot.Y = newY
}

func (r *Rope) MoveRight() {
	r.Head().X++
	r.onHeadMove()
}

func (r *Rope) MoveLeft() {
	r.Head().X--
	r.onHeadMove()
}

func (r *Rope) MoveUp() {
	r.Head().Y++
	r.onHeadMove()
}

func (r *Rope) MoveDown() {
	r.Head().Y--
	r.onHeadMove()
}
