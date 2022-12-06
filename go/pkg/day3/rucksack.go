package day3

import "unicode"

type Item rune

func (i Item) Priority() int {
	var offset int
	if unicode.IsUpper(rune(i)) {
		offset = 38
	} else {
		offset = 96
	}

	return int(i) - offset
}

type Rucksack struct {
	Items             []Item
	FirstCompartment  []Item
	SecondCompartment []Item
}

func NewRucksack(allItems []Item) *Rucksack {
	firstCompartment := allItems[0 : len(allItems)/2]
	secondCompartment := allItems[len(allItems)/2:]

	return &Rucksack{
		Items:             allItems,
		FirstCompartment:  firstCompartment,
		SecondCompartment: secondCompartment,
	}
}

func (r *Rucksack) FindDuplicateItem() (item Item, found bool) {
	items := intersect(r.FirstCompartment, r.SecondCompartment)
	if len(items) != 1 {
		return
	}

	item = items[0]
	found = true

	return
}

func intersect(a, b []Item) []Item {
	exists := struct{}{}
	intersection := make(map[Item]struct{}, 0)
	for _, aItem := range a {
		for _, bItem := range b {
			if aItem == bItem {
				intersection[aItem] = exists
			}
		}
	}

	list := make([]Item, 0)
	for value := range intersection {
		list = append(list, value)
	}

	return list
}

type Group struct {
	Rucksacks []Rucksack
}

func NewGroup(rucksacks []Rucksack) *Group {
	return &Group{Rucksacks: rucksacks}
}

func (g *Group) FindCommomItem() (item Item, found bool) {
	if len(g.Rucksacks) < 2 {
		return
	}

	intersection := intersect(g.Rucksacks[0].Items, g.Rucksacks[1].Items)
	for _, rucksack := range g.Rucksacks[2:] {
		intersection = intersect(intersection, rucksack.Items)
	}

	if len(intersection) == 1 {
		item = intersection[0]
		found = true
	}

	return item, found
}
