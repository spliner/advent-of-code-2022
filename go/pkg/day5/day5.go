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
	foo, commands, err := parseInput(input)
	if err != nil {
		return "", err
	}

	for _, command := range commands {
		command.execute(foo)
	}

	values := foo.peek()
	result := strings.Join(values, "")
	return result, nil
}

func parseInput(input string) (*foo, []command, error) {
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

	foo := &foo{orderedIds: orderedStackIds, stacks: stacks}
	// 5  -  2
	// 29 -  x
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
	commands := make([]command, 0)
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
		command := command{
			originStackId:      originStackId,
			destinationStackId: destinationStackId,
			quantity:           quantity,
		}
		commands = append(commands, command)
	}
	return foo, commands, nil
}

func Part2(input string) (string, error) {
	return "", nil
}
