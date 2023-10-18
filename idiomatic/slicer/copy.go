package slicer

func Copy[T any](items ...T) []T {
	copy := make([]T, len(items))

	for i := 0; i < len(items); i++ {
		copy[i] = items[i]
	}

	return copy
}
