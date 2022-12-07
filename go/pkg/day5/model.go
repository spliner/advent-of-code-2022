package day5

import (
	"aoc2022/pkg/util/stack"
	"fmt"
)

type Crane struct {
	orderedIds []string
	stacks     map[string]*stack.Stack[string]
}

func (c *Crane) Get(id string) (*stack.Stack[string], bool) {
	stack, ok := c.stacks[id]
	return stack, ok
}

func (c *Crane) Peek() []string {
	values := make([]string, len(c.orderedIds))
	for i, id := range c.orderedIds {
		value, _ := c.stacks[id].Peek()
		values[i] = value
	}
	return values
}

type Command struct {
	originStackId      string
	destinationStackId string
	quantity           int
}

func (c *Command) execute(crane *Crane) error {
	originStack, ok := crane.Get(c.originStackId)
	if !ok {
		return fmt.Errorf("stack %s not found", c.originStackId)
	}

	destinationStack, ok := crane.Get(c.destinationStackId)
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
