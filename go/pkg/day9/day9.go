package day9

import (
	"aoc2022/pkg/day9/model"
	"aoc2022/pkg/util/set"
	"fmt"
	"strconv"
	"strings"
)

var commandMap = map[string]func(r *model.Rope){
	"R": func(r *model.Rope) { r.MoveRight() },
	"U": func(r *model.Rope) { r.MoveUp() },
	"L": func(r *model.Rope) { r.MoveLeft() },
	"D": func(r *model.Rope) { r.MoveDown() },
}

func Part1(input string) (string, error) {
	count, err := execute(input)
	if err != nil {
		return "", err
	}

	result := strconv.Itoa(count)
	return result, nil
}

func execute(input string) (int, error) {
	set := set.NewSet[model.Position]()
	rope := model.NewRope()
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		split := strings.Split(line, " ")
		command, ok := commandMap[split[0]]
		if !ok {
			return 0, fmt.Errorf("could not find command %s", split[0])
		}

		repetitions, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, err
		}

		for i := 0; i < repetitions; i++ {
			command(rope)
			set.Add(*rope.Tail())
		}
	}

	return set.Length(), nil
}

func Part2(input string) (string, error) {
	return "", nil
}
