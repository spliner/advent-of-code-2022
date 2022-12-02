package util

func Map[T any, U any](source []T, mapper func(T) U) []U {
	s := make([]U, len(source))
	for i, value := range source {
		s[i] = mapper(value)
	}
	return s
}
