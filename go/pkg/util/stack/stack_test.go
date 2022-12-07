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

	got, ok := stack.Pop()
	if !ok {
		t.Fatal("expected Pop to return a value")
	}

	want := 3
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}

	got, ok = stack.Pop()
	if !ok {
		t.Fatal("expected Pop to return a value")
	}

	want = 2
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}

	got, ok = stack.Pop()
	if !ok {
		t.Fatal("expected Pop to return a value")
	}

	want = 1
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}

	_, ok = stack.Pop()
	if ok {
		t.Fatal("expected Pop to return no value")
	}
}
