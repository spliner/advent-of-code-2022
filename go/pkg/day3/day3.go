package day3

import (
	"errors"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var totalPriority int
	for _, line := range lines {
		items := parseLine(line)
		rucksack := NewRucksack(items)
		duplicateItem, found := rucksack.FindDuplicateItem()
		if !found {
			return "", errors.New("Rucksack has no duplicate item")
		}

		totalPriority += duplicateItem.Priority()
	}

	result := strconv.Itoa(totalPriority)
	return result, nil
}

func parseLine(line string) []Item {
	return []Item(line)
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")

	var totalPriority int
	for i := 0; i < len(lines); i += 3 {
		rucksacks := make([]Rucksack, 3)
		for j, line := range lines[i : i+3] {
			items := parseLine(line)
			rucksack := NewRucksack(items)
			rucksacks[j] = *rucksack
		}

		group := NewGroup(rucksacks)
		commomItem, found := group.FindCommomItem()
		if !found {
			return "", errors.New("Group has no item in commom")
		}

		totalPriority += commomItem.Priority()
	}

	result := strconv.Itoa(totalPriority)
	return result, nil
}
