package slices

func All[T comparable](xs []T, v T) []T {
	return AllFn[T](xs, v, func(x, y T) bool { return x == y })
}

func AllFn[T comparable](xs []T, v T, fn func(T, T) bool) []T {
	var os []T

	for _, x := range xs {
		if fn(x, v) {
			os = append(os, x)
		}
	}

	return os
}