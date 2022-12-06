package day3_test

import (
	"aoc2022/pkg/day3"
	"reflect"
	"testing"
)

func TestItemPriority(t *testing.T) {
	t.Run("'p' should be 16", func(t *testing.T) {
		var item day3.Item = 'p'
		got := item.Priority()
		want := 16
		if got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	})

	t.Run("'L' should be 38", func(t *testing.T) {
		var item day3.Item = 'L'
		got := item.Priority()
		want := 38
		if got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	})
}

func TestRuckSackCompartments(t *testing.T) {
	items := []day3.Item{'a', 'b', 'c', 'd'}

	rucksack := day3.NewRucksack(items)

	expectedFirstCompartment := []day3.Item{'a', 'b'}
	if !reflect.DeepEqual(rucksack.FirstCompartment, expectedFirstCompartment) {
		t.Fatalf("got %v, expected %v", rucksack.FirstCompartment, expectedFirstCompartment)
	}

	expectedSecondCompartment := []day3.Item{'c', 'd'}
	if !reflect.DeepEqual(rucksack.SecondCompartment, expectedSecondCompartment) {
		t.Fatalf("got %v, expected %v", rucksack.SecondCompartment, expectedSecondCompartment)
	}
}

func TestPart1(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

	got, err := day3.Part1(input)
	if err != nil {
		t.Fatal("expected no error", err)
	}

	want := "157"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestPart2(t *testing.T) {
}
