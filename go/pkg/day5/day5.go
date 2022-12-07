package day5

import (
	"aoc2022/pkg/util/stack"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Part1(input string) (string, error) {
	crane, commands, err := parseInput(input)
	if err != nil {
		return "", err
	}

	for _, command := range commands {
		command.execute(crane)
	}

	values := crane.Peek()
	result := strings.Join(values, "")
	return result, nil
}

func parseInput(input string) (*Crane, []Command, error) {
	split := strings.Split(input, "\n\n")
	stackLines := strings.Split(split[0], "\n")

	stacks := make(map[string]*stack.Stack[string])
	stackIdsLine := stackLines[len(stackLines)-1]
	stackIds := strings.Split(stackIdsLine, " ")
	orderedStackIds := make([]string, 0)
	for _, stackId := range stackIds {
		stackId = strings.Trim(stackId, " ")
		if stackId != "" {
			orderedStackIds = append(orderedStackIds, stackId)
			stacks[stackId] = stack.NewStack[string]()
		}
	}

	crane := &Crane{orderedIds: orderedStackIds, stacks: stacks}
	for i := len(stackLines) - 2; i >= 0; i-- {
		line := stackLines[i]
		runes := []rune(line)
		for j, r := range runes {
			if unicode.IsLetter(r) {
				stackId := strconv.Itoa((j-1)/4 + 1)
				stacks[stackId].Push(string(r))
			}
		}
	}

	re, err := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)
	if err != nil {
		return nil, nil, err
	}

	commandLines := strings.Split(split[1], "\n")
	commands := make([]Command, 0)
	for _, line := range commandLines {
		matches := re.FindStringSubmatch(line)
		if len(matches) != 4 {
			return nil, nil, fmt.Errorf("expected 4 matches, got %d", len(matches))
		}
		quantity, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, nil, err
		}

		originStackId := matches[2]
		destinationStackId := matches[3]
		command := Command{
			originStackId:      originStackId,
			destinationStackId: destinationStackId,
			quantity:           quantity,
		}
		commands = append(commands, command)
	}
	return crane, commands, nil
}

func Part2(input string) (string, error) {
	return "", nil
}
