package day6_test

import (
	"aoc2022/pkg/day6"
	"fmt"
	"testing"
)

type part1TestCase struct {
	input       string
	part1Marker string
	part2Marker string
}

var testCases = []part1TestCase{
	{
		input:       "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		part1Marker: "7",
		part2Marker: "19",
	},
	{
		input:       "bvwbjplbgvbhsrlpgdmjqwftvncz",
		part1Marker: "5",
		part2Marker: "23",
	},
	{
		input:       "nppdvjthqldpwncqszvftbrmjlhg",
		part1Marker: "6",
		part2Marker: "23",
	},
}

func TestPart1(t *testing.T) {
	for _, testCase := range testCases {
		name := fmt.Sprintf("%s should return %ss", testCase.input, testCase.part1Marker)
		t.Run(name, func(t *testing.T) {
			got, err := day6.Part1(testCase.input)
			if err != nil {
				t.Fatal("expected no error, got", err)
			}

			if got != testCase.part1Marker {
				t.Fatalf("got %s, want %s", got, testCase.part1Marker)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	for _, testCase := range testCases {
		name := fmt.Sprintf("%s should return %ss", testCase.input, testCase.part2Marker)
		t.Run(name, func(t *testing.T) {
			got, err := day6.Part2(testCase.input)
			if err != nil {
				t.Fatal("expected no error, got", err)
			}

			if got != testCase.part2Marker {
				t.Fatalf("got %s, want %s", got, testCase.part2Marker)
			}
		})
	}
}
