package day8_test

import (
	"aoc2022/pkg/day8"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `30373
25512
65332
33549
35390`

	got, err := day8.Part1(input)
	if err != nil {
		t.Fatal("expected no error, got ", err)
	}

	want := "21"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `30373
25512
65332
33549
35390`

	got, err := day8.Part2(input)
	if err != nil {
		t.Fatal("expected no error, got ", err)
	}

	want := "8"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
