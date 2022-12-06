package day4

type Range struct {
	Start int
	End   int
}

func NewRange(start, end int) *Range {
	return &Range{Start: start, End: end}
}

func (r *Range) Size() int {
	return r.End - r.Start
}

type AssignmentPair struct {
	FirstRange  *Range
	SecondRange *Range
}

func NewAssignmentPair(firstRange, secondRange *Range) *AssignmentPair {
	return &AssignmentPair{FirstRange: firstRange, SecondRange: secondRange}
}

func (p *AssignmentPair) HasFullOverlap() bool {
	var smallerRange, largerRange *Range
	if p.FirstRange.Size() < p.SecondRange.Size() {
		smallerRange = p.FirstRange
		largerRange = p.SecondRange
	} else {
		smallerRange = p.SecondRange
		largerRange = p.FirstRange
	}

	return largerRange.Start <= smallerRange.Start && largerRange.End >= smallerRange.End
}
