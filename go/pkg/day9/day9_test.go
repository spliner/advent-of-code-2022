package day9_test

import (
	"aoc2022/pkg/day9"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	got, err := day9.Part1(input)
	if err != nil {
		t.Fatal("expected no error, got ", err)
	}

	want := "13"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	got, err := day9.Part2(input)
	if err != nil {
		t.Fatal("expected no error, got ", err)
	}

	want := "1"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestPart2Input2(t *testing.T) {
	input := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	got, _ := day9.Part2(input)

	want := "36"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
