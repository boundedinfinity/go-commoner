package slicer

func Contains[T comparable](xs []T, v T) bool {
	fn := func(x T) bool {
		return x == v
	}

	return ContainsFn(xs, fn)
}

func ContainsFn[T comparable](xs []T, fn func(T) bool) bool {
	for _, x := range xs {
		if fn(x) {
			return true
		}
	}

	return false
}
