package slices

func Find[T comparable](ss []T, t T) (T, bool) {
	fn := func(s T) bool {
		return s == t
	}

	return FindFn(ss, fn)
}

func FindFn[T any](ss []T, fn func(T) bool) (T, bool) {
	for _, s := range ss {
		if fn(s) {
			return s, true
		}
	}

	var zero T
	return zero, false
}
