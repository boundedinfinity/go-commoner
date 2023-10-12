package slicer

func Chunk[T comparable](size int, items ...T) [][]T {
	os := [][]T{}
	cur := []T{}

	for i := 0; i < len(items); i++ {
		if i%size == 0 && i != 0 {
			os = append(os, cur)
			cur = []T{}
		}

		cur = append(cur, items[i])
	}

	if len(cur) > 0 {
		os = append(os, cur)
	}

	return os
}
