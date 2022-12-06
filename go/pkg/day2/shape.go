package day2

type shapeKind int

const (
	rockKind shapeKind = iota
	paperKind
	scissorsKind
)

type shape struct {
	value int
	kind  shapeKind
}

var rock = &shape{1, rockKind}
var paper = &shape{2, paperKind}
var scissors = &shape{3, scissorsKind}

var winMap = map[shapeKind]*shape{
	rockKind:     scissors,
	paperKind:    rock,
	scissorsKind: paper,
}

var lossMap = map[shapeKind]*shape{
	rockKind:     paper,
	paperKind:    scissors,
	scissorsKind: rock,
}

func (s *shape) outcome(other *shape) *outcome {
	if s.kind == other.kind {
		return drawOutcome
	}

	if winMap[s.kind].kind == other.kind {
		return winOutcome
	}

	return lossOutcome
}

func (s *shape) shapeToPlay(outcome *outcome) *shape {
	if outcome.kind == draw {
		return s
	}

	if outcome.kind == win {
		return lossMap[s.kind]
	}

	return winMap[s.kind]
}
