package slicer

func Find[T comparable](match T, elems ...T) (T, bool) {
	fn := func(_ int, elem T) bool {
		return elem == match
	}

	return FindFn(fn, elems...)
}

func FindFn[T any](fn func(int, T) bool, elems ...T) (T, bool) {
	if fn == nil {
		var zero T
		return zero, false
	}

	fn2 := func(i int, elem T) (bool, error) {
		return fn(i, elem), nil
	}

	found, ok, _ := FindFnErr(fn2, elems...)
	return found, ok
}

func FindFnErr[T any](fn func(int, T) (bool, error), elems ...T) (T, bool, error) {
	var found T
	var ok bool
	var err error

	if fn == nil {
		return found, ok, err
	}

	for i, elem := range elems {
		if ok, err = fn(i, elem); ok || err != nil {
			found = elem
			break
		}
	}

	return found, ok, err
}
