package slices

func Chunk[T comparable](xs []T, size int) [][]T {
	os := [][]T{}
	cur := []T{}

	for i := 0; i < len(xs); i++ {
		if i%size == 0 && i != 0 {
			os = append(os, cur)
			cur = []T{}
		}

		cur = append(cur, xs[i])
	}

	if len(cur) > 0 {
		os = append(os, cur)
	}

	return os
}
