package day6

import (
	"aoc2022/pkg/util/set"
	"strconv"
)

func Part1(input string) (string, error) {
	marker := marker(input, 4)
	result := strconv.Itoa(marker)
	return result, nil
}

func marker(input string, size int) int {
	runes := []rune(input)
	for i := range runes {
		slice := runes[i : i+size]
		set := set.NewSetFromSlice(slice)
		if set.Length() == len(slice) {
			return i + size
		}
	}

	return 0
}

func Part2(input string) (string, error) {
	return "", nil
}
