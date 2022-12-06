package day4

import (
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var count int
	for _, line := range lines {
		pair, err := parseLine(line)
		if err != nil {
			return "", err
		}

		if pair.HasFullOverlap() {
			count++
		}
	}

	result := strconv.Itoa(count)
	return result, nil
}

func parseLine(line string) (*AssignmentPair, error) {
	rawRanges := strings.Split(line, ",")
	firstRange, err := parseRange(rawRanges[0])
	if err != nil {
		return nil, err
	}

	secondRange, err := parseRange(rawRanges[1])
	if err != nil {
		return nil, err
	}

	return NewAssignmentPair(firstRange, secondRange), nil
}

func parseRange(rawRange string) (*Range, error) {
	split := strings.Split(rawRange, "-")
	start, err := strconv.Atoi(split[0])
	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(split[1])
	if err != nil {
		return nil, err
	}

	return NewRange(start, end), nil
}

func Part2(input string) (string, error) {
	return "", nil
}
