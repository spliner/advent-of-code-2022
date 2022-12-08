package day6_test

import (
	"aoc2022/pkg/day6"
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	type part1TestCase struct {
		input       string
		firstMarker string
	}
	testCases := []part1TestCase{
		{
			input:       "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			firstMarker: "7",
		},
		{
			input:       "bvwbjplbgvbhsrlpgdmjqwftvncz",
			firstMarker: "5",
		},
		{
			input:       "nppdvjthqldpwncqszvftbrmjlhg",
			firstMarker: "6",
		},
	}
	for _, testCase := range testCases {
		name := fmt.Sprintf("%s should return %ss", testCase.input, testCase.firstMarker)
		t.Run(name, func(t *testing.T) {
			got, err := day6.Part1(testCase.input)
			if err != nil {
				t.Fatal("expected no error, got", err)
			}

			if got != testCase.firstMarker {
				t.Fatalf("got %s, want %s", got, testCase.firstMarker)
			}
		})
	}
}
