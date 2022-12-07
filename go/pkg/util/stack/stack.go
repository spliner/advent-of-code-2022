package stack

type Stack[T any] struct {
	head *node[T]
}

type node[T any] struct {
	value    *T
	previous *node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	node := &node[T]{value: &value, previous: s.head}
	s.head = node
}

func (s *Stack[T]) Peek() (value T, ok bool) {
	if s.head == nil {
		return
	}

	value = *s.head.value
	ok = true

	return value, ok
}

func (s *Stack[T]) Pop() (value T, ok bool) {
	if s.head == nil {
		return
	}

	value = *s.head.value
	ok = true

	s.head = s.head.previous

	return
}
