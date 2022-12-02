package util

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Sum[T Number](items []T) T {
	var sum T
	for _, value := range items {
		sum += value
	}
	return sum
}
