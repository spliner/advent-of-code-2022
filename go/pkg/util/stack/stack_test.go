package stack_test

import (
	"aoc2022/pkg/util/stack"
	"testing"
)

func TestPush(t *testing.T) {
	stack := stack.NewStack[string]()
	stack.Push("a")
	stack.Push("b")

	got, ok := stack.Peek()
	if !ok {
		t.Fatal("expected Peek to return a value")
	}

	want := "b"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestPeekWithEmptyStack(t *testing.T) {
	stack := stack.NewStack[bool]()

	_, got := stack.Peek()

	want := false
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestPop(t *testing.T) {
	stack := stack.NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	assert := func(expectedValue, actualValue int, ok bool) {
		if !ok {
			t.Fatal("expected Pop to return a value")
		}

		if actualValue != expectedValue {
			t.Fatalf("got %d, want %d", actualValue, expectedValue)
		}
	}

	got, ok := stack.Pop()
	assert(3, got, ok)

	got, ok = stack.Pop()
	assert(2, got, ok)

	got, ok = stack.Pop()
	assert(1, got, ok)

	_, ok = stack.Pop()
	if ok {
		t.Fatal("expected Pop to return no value")
	}
}
