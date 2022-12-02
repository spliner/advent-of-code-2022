package day1

import (
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	elves, err := parseInput(input)
	if err != nil {
		return "", err
	}

	mostCaloriesElf := elves[0]

	result := strconv.Itoa(mostCaloriesElf.TotalCalories())
	return result, nil
}

func parseInput(input string) ([]elf, error) {
	split := strings.Split(input, "\n\n")

	elves := make([]elf, len(split))
	for i, lines := range split {
		split = strings.Split(lines, "\n")

		meals := make([]meal, len(split))
		for j, line := range split {
			calories, err := strconv.Atoi(line)
			if err != nil {
				return nil, err
			}

			meals[j] = meal{calories: calories}
		}

		elves[i] = elf{meals: meals}
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].TotalCalories() > elves[j].TotalCalories()
	})

	return elves, nil
}

func Part2(input string) (string, error) {
	elves, err := parseInput(input)
	if err != nil {
		return "", err
	}

	var total int
	for i := 0; i < 3; i++ {
		total += elves[i].TotalCalories()
	}

	result := strconv.Itoa(total)
	return result, nil
}
