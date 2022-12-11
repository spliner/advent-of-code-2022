package day10

import (
	"aoc2022/pkg/day10/model"
	"aoc2022/pkg/util/set"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	cyclesToLog := set.NewSet[int]()
	cyclesToLog.Add(20)
	cyclesToLog.Add(60)
	cyclesToLog.Add(100)
	cyclesToLog.Add(140)
	cyclesToLog.Add(180)
	cyclesToLog.Add(220)

	cpu := model.NewCpu(cyclesToLog)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		instruction, err := parseInstruction(line)
		if err != nil {
			return "", err
		}

		cpu.Execute(instruction)
	}

	result := strconv.Itoa(cpu.Sum())
	return result, nil
}

func parseInstruction(rawInstruction string) (model.Instruction, error) {
	split := strings.Split(rawInstruction, " ")
	if split[0] == "noop" {
		return model.NewNoopInstruction(), nil
	}

	if split[0] == "addx" {
		value, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}

		return model.NewAddInstruction(value), nil
	}

	return nil, fmt.Errorf("unknown instruction: %s", rawInstruction)
}

func Part2(input string) (string, error) {
	return "", nil
}
