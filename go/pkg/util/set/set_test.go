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

func TestRemoveExistingItem(t *testing.T) {
	set := set.NewSet[string]()
	set.Add("test")

	got := set.Remove("test")

	want := true
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	if !set.IsEmpty() {
		t.Fatal("expected set to be empty")
	}
}

func TestRemoveNonExistingItem(t *testing.T) {
	set := set.NewSet[string]()
	set.Add("test")

	got := set.Remove("foo")

	want := false
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	if set.IsEmpty() {
		t.Fatal("expected set to not be empty")
	}
}

func TestLenght(t *testing.T) {
	set := set.NewSet[string]()

	set.Add("1")
	if set.Length() != 1 {
		t.Fatal("expected length to be 1")
	}

	set.Add("2")
	if set.Length() != 2 {
		t.Fatal("expected length to be 2")
	}

	set.Remove("1")
	if set.Length() != 1 {
		t.Fatal("expected length to be 1")
	}
}
