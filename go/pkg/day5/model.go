package day5

import (
	"aoc2022/pkg/util/stack"
	"fmt"
)

type foo struct {
	orderedIds []string
	stacks     map[string]*stack.Stack[string]
}

func (s *foo) get(id string) (*stack.Stack[string], bool) {
	stack, ok := s.stacks[id]
	return stack, ok
}

func (s *foo) peek() []string {
	values := make([]string, len(s.orderedIds))
	for i, id := range s.orderedIds {
		value, _ := s.stacks[id].Peek()
		values[i] = value
	}
	return values
}

type command struct {
	originStackId      string
	destinationStackId string
	quantity           int
}

func (c *command) execute(stacks *foo) error {
	originStack, ok := stacks.get(c.originStackId)
	if !ok {
		return fmt.Errorf("stack %s not found", c.originStackId)
	}

	destinationStack, ok := stacks.get(c.destinationStackId)
	if !ok {
		return fmt.Errorf("stack %s not found", c.destinationStackId)
	}

	for i := 0; i < c.quantity; i++ {
		create, ok := originStack.Pop()
		if !ok {
			break
		}

		destinationStack.Push(create)
	}

	return nil
}
