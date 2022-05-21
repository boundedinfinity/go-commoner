package slices

func Filter[T any](xs []T, fn func(T) bool) []T {
	var os []T

	for _, x := range xs {
		if fn(x) {
			os = append(os, x)
		}
	}

	return os
}
