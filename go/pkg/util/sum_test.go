package util_test

import (
	"aoc2022/pkg/util"
	"testing"
)

func TestSumInts(t *testing.T) {
	items := []int{1, 2, 3}

	got := util.Sum(items)

	want := 6
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestSumFloats(t *testing.T) {
	items := []float64{1, 2, 3}

	got := util.Sum(items)

	want := 6.0
	if got != want {
		t.Fatalf("got %f, want %f", got, want)
	}
}
