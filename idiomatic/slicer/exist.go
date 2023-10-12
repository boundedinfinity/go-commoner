package slicer

func Exist[T comparable](s T, items ...T) bool {
	fn := func(t T) bool {
		return t == s
	}

	return ExistFn(fn, items...)
}

func ExistFn[T any](fn func(T) bool, items ...T) bool {
	for _, item := range items {
		if fn(item) {
			return true
		}
	}

	return false
}

func ExistFnErr[T any](fn func(T) (bool, error), items ...T) (bool, error) {
	var ok bool
	var err error

	for _, item := range items {
		if ok, err = fn(item); ok || err != nil {
			break
		}
	}

	return ok, err
}
