package model

type Instruction interface {
	NumberOfCycles() int
	Execute(currentRegister int) (newRegister int)
}

type NoopInstruction struct {
}

func NewNoopInstruction() *NoopInstruction {
	return &NoopInstruction{}
}

func (i *NoopInstruction) NumberOfCycles() int {
	return 1
}

func (i *NoopInstruction) Execute(currentRegister int) (newRegister int) {
	return currentRegister
}

type AddInstruction struct {
	value int
}

func NewAddInstruction(value int) *AddInstruction {
	return &AddInstruction{value: value}
}

func (i *AddInstruction) NumberOfCycles() int {
	return 2
}

func (i *AddInstruction) Execute(currentRegister int) (newRegister int) {
	return currentRegister + i.value
}
