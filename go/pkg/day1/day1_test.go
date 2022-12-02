package day1_test

import (
	"aoc2022/pkg/day1"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	got, err := day1.Part1(input)
	if err != nil {
		t.Fatal("expected no error", err)
	}

	want := "24000"

	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestDay1Part2(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	got, err := day1.Part2(input)
	if err != nil {
		t.Fatal("expected no error", err)
	}

	want := "45000"

	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
