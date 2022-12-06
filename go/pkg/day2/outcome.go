package day2

type outcome struct {
	value int
	kind  outcomeKind
}

type outcomeKind int

const (
	win outcomeKind = iota
	draw
	loss
)

var winOutcome = &outcome{6, win}
var drawOutcome = &outcome{3, draw}
var lossOutcome = &outcome{0, loss}
