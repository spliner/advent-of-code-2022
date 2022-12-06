package day2_test

import (
	"aoc2022/pkg/day2"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	input := `A Y
B X
C Z`

	got, err := day2.Part1(input)
	if err != nil {
		t.Fatal("expected no error", err)
	}

	want := "15"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestDay2Part2(t *testing.T) {
	input := `A Y
B X
C Z`

	got, err := day2.Part2(input)
	if err != nil {
		t.Fatal("expected no error", err)
	}

	want := "12"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
