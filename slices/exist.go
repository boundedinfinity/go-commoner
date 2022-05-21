package slices

func Exist[T comparable](ss []T, s T) bool {
	fn := func(t T) bool {
		return t == s
	}

	return ExistFn(ss, fn)
}

func ExistFn[T any](ss []T, fn func(T) bool) bool {
	for _, s := range ss {
		if fn(s) {
			return true
		}
	}

	return false
}
