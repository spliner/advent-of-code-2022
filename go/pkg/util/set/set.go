package set

var marker struct{}

type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]struct{})}
}

func NewSetFromSlice[T comparable](slice []T) *Set[T] {
	set := NewSet[T]()
	for _, item := range slice {
		set.Add(item)
	}
	return set
}

func (s *Set[T]) Length() int {
	return len(s.items)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Set[T]) Add(item T) bool {
	if _, ok := s.items[item]; ok {
		return false
	}

	s.items[item] = marker
	return true
}

func (s *Set[T]) Remove(item T) bool {
	if _, ok := s.items[item]; !ok {
		return false
	}

	delete(s.items, item)
	return true
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}
