package slicer

func Find[T comparable](match T, items ...T) (T, bool) {
	fn := func(item T) bool {
		return item == match
	}

	return FindFn(fn, items...)
}

func FindFn[T any](fn func(T) bool, items ...T) (T, bool) {
	var ok bool
	var item T

	for _, item = range items {
		if ok = fn(item); ok {
			break
		}
	}

	return item, ok
}

func FindFnErr[T any](fn func(T) (bool, error), items ...T) (T, bool, error) {
	var item T
	var ok bool
	var err error

	for _, item = range items {
		if ok, err = fn(item); ok || err != nil {
			break
		}
	}

	return item, ok, err
}
