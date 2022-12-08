package set

var marker struct{}

type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]struct{})}
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Set[T]) Add(item T) bool {
	if _, ok := s.items[item]; ok {
		return false
	}

	s.items[item] = marker
	return true
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}
