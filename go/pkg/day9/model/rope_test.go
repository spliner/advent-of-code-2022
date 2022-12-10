package model_test

import (
	"aoc2022/pkg/day9/model"
	"fmt"
	"testing"
)

type linearMoveTestCase struct {
	moveFunc     func(r *model.Rope)
	direction    string
	expectedHead model.Position
	expectedTail model.Position
}

func TestLinearMove(t *testing.T) {
	testCases := []linearMoveTestCase{
		{
			moveFunc:     func(r *model.Rope) { r.MoveRight() },
			direction:    "right",
			expectedHead: model.Position{2, 0},
			expectedTail: model.Position{1, 0},
		},
		{
			moveFunc:     func(r *model.Rope) { r.MoveLeft() },
			direction:    "left",
			expectedHead: model.Position{-2, 0},
			expectedTail: model.Position{-1, 0},
		},
		{
			moveFunc:     func(r *model.Rope) { r.MoveUp() },
			direction:    "up",
			expectedHead: model.Position{0, 2},
			expectedTail: model.Position{0, 1},
		},
		{
			moveFunc:     func(r *model.Rope) { r.MoveDown() },
			direction:    "down",
			expectedHead: model.Position{0, -2},
			expectedTail: model.Position{0, -1},
		},
	}

	for _, test := range testCases {
		name := fmt.Sprintf("test move %s", test.direction)
		t.Run(name, func(t *testing.T) {
			rope := model.NewRope(2)
			test.moveFunc(rope)
			test.moveFunc(rope)

			got := rope.Head()
			if *got != test.expectedHead {
				t.Fatalf("got %v, want %v", *got, test.expectedHead)
			}

			got = rope.Tail()
			if *got != test.expectedTail {
				t.Fatalf("got %v, want %v", *got, test.expectedTail)
			}
		})
	}
}

func TestDiagonalMove(t *testing.T) {
	rope := model.NewRope(2)
	rope.MoveRight()
	rope.MoveUp()
	rope.MoveUp()

	got := rope.Head()

	want := model.Position{1, 2}
	if *got != want {
		t.Fatalf("got %v, want %v", *got, want)
	}

	got = rope.Tail()

	want = model.Position{1, 1}
	if *got != want {
		t.Fatalf("got %v, want %v", *got, want)
	}
}
