package day5_test

import (
	"aoc2022/pkg/day5"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
	1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	got, err := day5.Part1(input)
	if err != nil {
		t.Fatal("expected no error, got", err)
	}

	want := "CMZ"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
	1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	got, err := day5.Part2(input)
	if err != nil {
		t.Fatal("expected no error, got", err)
	}

	want := "TODO"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
