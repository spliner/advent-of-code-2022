package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var total int
	for _, line := range lines {
		myShape, opponentShape, err := parseShapes(line)
		if err != nil {
			return "", err
		}

		round := round{myShape: myShape, opponentShape: opponentShape}
		total += round.myScore()
	}

	return strconv.Itoa(total), nil
}

func parseShapes(line string) (myShape *shape, opponentShape *shape, err error) {
	rawShapes := strings.Split(line, " ")
	opponentShape, err = parseShape(rawShapes[0])
	if err != nil {
		return nil, nil, err
	}

	myShape, err = parseShape(rawShapes[1])
	if err != nil {
		return nil, nil, err
	}

	return
}

func parseShape(rawShape string) (*shape, error) {
	switch rawShape {
	case "A", "X":
		return rock, nil
	case "B", "Y":
		return paper, nil
	case "C", "Z":
		return scissors, nil
	default:
		return nil, fmt.Errorf("invalid raw shape: %s", rawShape)
	}
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var total int
	for _, line := range lines {
		opponentShape, desiredOutcome, err := foo(line)
		if err != nil {
			return "", err
		}

		myShape := opponentShape.shapeToPlay(desiredOutcome)
		round := round{myShape: myShape, opponentShape: opponentShape}
		total += round.myScore()
	}

	return strconv.Itoa(total), nil
}

func foo(line string) (opponentShape *shape, desiredOutcome *outcome, err error) {
	split := strings.Split(line, " ")
	opponentShape, err = parseShape(split[0])
	if err != nil {
		return nil, nil, err
	}

	desiredOutcome, err = parseOutcome(split[1])
	if err != nil {
		return nil, nil, err
	}

	return
}

func parseOutcome(rawOutcome string) (*outcome, error) {
	switch rawOutcome {
	case "X":
		return lossOutcome, nil
	case "Y":
		return drawOutcome, nil
	case "Z":
		return winOutcome, nil
	default:
		return nil, fmt.Errorf("invalid raw outcome: %s", rawOutcome)
	}
}
