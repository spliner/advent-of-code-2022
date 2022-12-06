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
	FirstCompartment  []Item
	SecondCompartment []Item
}

func NewRucksack(allItems []Item) *Rucksack {
	firstCompartment := allItems[0 : len(allItems)/2]
	secondCompartment := allItems[len(allItems)/2:]

	return &Rucksack{
		FirstCompartment:  firstCompartment,
		SecondCompartment: secondCompartment,
	}
}

func (r *Rucksack) FindDuplicateItem() (item Item, found bool) {
	for i := 0; i < len(r.FirstCompartment); i++ {
		firstItem := r.FirstCompartment[i]
		for j := 0; j < len(r.SecondCompartment); j++ {
			secondItem := r.SecondCompartment[j]
			if firstItem == secondItem {
				item = firstItem
				found = true
				break
			}
		}
	}

	return item, found
}
