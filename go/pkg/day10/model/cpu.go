package model

import "aoc2022/pkg/util/set"

type Cpu struct {
	cycle       int
	registerX   int
	sum         int
	cyclesToLog *set.Set[int]
}

func NewCpu(cyclesToLog *set.Set[int]) *Cpu {
	return &Cpu{
		cycle:       0,
		registerX:   1,
		sum:         0,
		cyclesToLog: cyclesToLog,
	}
}

func (c *Cpu) Execute(instruction Instruction) {
	for i := 0; i < instruction.NumberOfCycles()-1; i++ {
		c.cycle++
		if c.cyclesToLog.Contains(c.cycle) {
			c.sum += c.signalStrength()
		}
	}

	c.cycle++
	if c.cyclesToLog.Contains(c.cycle) {
		c.sum += c.signalStrength()
	}

	c.registerX = instruction.Execute(c.registerX)
}

func (c *Cpu) signalStrength() int {
	return c.cycle * c.registerX
}

func (c *Cpu) Sum() int {
	return c.sum
}
