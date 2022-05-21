package slices

func Contains[T comparable](xs []T, v T) bool {
	fn := func(x, y T) bool {
		return x == y
	}

	return ContainsFn(xs, v, fn)
}

func ContainsFn[T comparable](xs []T, v T, fn func(T, T) bool) bool {
	for _, x := range xs {
		if fn(x, v) {
			return true
		}
	}

	return false
}
