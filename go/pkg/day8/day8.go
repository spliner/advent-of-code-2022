package day8

import (
	"aoc2022/pkg/day8/model"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	patch := parseInput(input)

	var visibleTrees int
	for _, trees := range patch.Trees() {
		for _, tree := range trees {
			if patch.IsVisible(tree) {
				visibleTrees++
			}
		}
	}

	result := strconv.Itoa(visibleTrees)
	return result, nil
}

func parseInput(input string) *model.Patch {
	lines := strings.Split(input, "\n")
	trees := make([][]*model.Tree, len(lines))
	for y, line := range lines {
		for x, r := range line {
			height := int(r - '0')
			tree := model.NewTree(height, x, y)
			trees[y] = append(trees[y], tree)
		}
	}

	return model.NewPatch(len(lines), len(lines), trees)
}

func Part2(input string) (string, error) {
	patch := parseInput(input)
	var maxScore int
	for _, trees := range patch.Trees() {
		for _, tree := range trees {
			score := patch.ScenicScore(tree)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	result := strconv.Itoa(maxScore)
	return result, nil
}
