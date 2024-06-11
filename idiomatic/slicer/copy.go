package slicer

func Copy[T any](elems ...T) []T {
	copy := make([]T, len(elems))

	for i := 0; i < len(elems); i++ {
		copy[i] = elems[i]
	}

	return copy
}
