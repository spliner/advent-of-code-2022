package set_test

import (
	"aoc2022/pkg/util/set"
	"testing"
)

func TestNewSet(t *testing.T) {
	set := set.NewSet[string]()

	got := set.IsEmpty()

	want := true
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestAdd(t *testing.T) {
	set := set.NewSet[string]()

	got := set.Add("test")

	want := true
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	got = set.Add("test")

	want = false
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestContains(t *testing.T) {
	set := set.NewSet[int]()
	set.Add(10)

	t.Run("should return true if item is present", func(t *testing.T) {
		got := set.Contains(10)

		want := true
		if got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("should return false if item is not present", func(t *testing.T) {
		got := set.Contains(11)

		want := false
		if got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

}
