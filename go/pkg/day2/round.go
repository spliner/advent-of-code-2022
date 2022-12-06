package day2

type round struct {
	myShape       *shape
	opponentShape *shape
}

func (r *round) myScore() int {
	outcome := r.myShape.outcome(r.opponentShape)
	return r.myShape.value + outcome.value
}
