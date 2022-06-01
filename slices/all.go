package slices

func All[T comparable](xs []T, v T) []T {
	fn := func(x T) bool {
		return v == x
	}

	return AllFn(xs, fn)
}

func AllFn[T comparable](xs []T, fn func(T) bool) []T {
	var os []T

	for _, x := range xs {
		if fn(x) {
			os = append(os, x)
		}
	}

	return os
}
