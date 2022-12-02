package util_test

import (
	"aoc2022/pkg/util"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	source := []int{1, 2, 3}

	got := util.Map(source, func(v int) int { return v * 2 })

	want := []int{2, 4, 6}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
