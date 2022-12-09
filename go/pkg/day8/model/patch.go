package model

type Patch struct {
	height int
	width  int
	trees  [][]*Tree
}

func NewPatch(height, width int, trees [][]*Tree) *Patch {
	return &Patch{
		height: height,
		width:  width,
		trees:  trees,
	}
}

func (p *Patch) Trees() [][]*Tree {
	return p.trees
}

func (p *Patch) IsVisible(tree *Tree) bool {
	isEdge := tree.X() == 0 ||
		tree.Y() == 0 ||
		tree.X() == p.width-1 ||
		tree.Y() == p.height-1
	if isEdge {
		return isEdge
	}

	visible := true
	// Check if visible from the left
	for x := 0; x < tree.X(); x++ {
		otherTree := p.trees[tree.Y()][x]
		if otherTree.Height() >= tree.Height() {
			visible = false
			break
		}
	}

	if visible {
		return visible
	}

	visible = true
	// Check if visible from the right
	for x := p.width - 1; x > tree.X(); x-- {
		otherTree := p.trees[tree.Y()][x]
		if otherTree.Height() >= tree.Height() {
			visible = false
			break
		}
	}

	if visible {
		return visible
	}

	visible = true
	// Check if visible from the top
	for y := 0; y < tree.Y(); y++ {
		otherTree := p.trees[y][tree.X()]
		if otherTree.Height() >= tree.Height() {
			visible = false
			break
		}
	}

	if visible {
		return visible
	}

	visible = true
	// Check if visible from the bottom
	for y := p.height - 1; y > tree.Y(); y-- {
		otherTree := p.trees[y][tree.X()]
		if otherTree.Height() >= tree.Height() {
			visible = false
			break
		}
	}

	return visible
}

// I'm pretty sure I can use this to calculate whether
// a Tree is visible or not, but I'm already late in
// AoC's calendar!
func (p *Patch) ScenicScore(tree *Tree) int {
	leftScore := 0
	for x := tree.X() - 1; x >= 0; x-- {
		leftScore++
		otherTree := p.trees[tree.Y()][x]
		if otherTree.Height() >= tree.Height() {
			break
		}
	}

	rightScore := 0
	for x := tree.X() + 1; x < p.width; x++ {
		rightScore++
		otherTree := p.trees[tree.Y()][x]
		if otherTree.Height() >= tree.Height() {
			break
		}
	}

	topScore := 0
	for y := tree.Y() - 1; y >= 0; y-- {
		topScore++
		otherTree := p.trees[y][tree.X()]
		if otherTree.Height() >= tree.Height() {
			break
		}
	}

	bottomScore := 0
	for y := tree.Y() + 1; y < p.height; y++ {
		bottomScore++
		otherTree := p.trees[y][tree.X()]
		if otherTree.Height() >= tree.Height() {
			break
		}
	}

	return leftScore *
		rightScore *
		topScore *
		bottomScore
}
