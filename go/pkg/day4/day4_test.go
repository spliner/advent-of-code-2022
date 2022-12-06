package day4_test

import (
	"aoc2022/pkg/day4"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	got, err := day4.Part1(input)
	if err != nil {
		t.Fatal("expected no error, got", err)
	}

	want := "2"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	got, err := day4.Part2(input)
	if err != nil {
		t.Fatal("expected no error, got", err)
	}

	want := "TODO"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
