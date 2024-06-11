package slicer

func Find[T comparable](match T, elems ...T) (T, bool) {
	fn := func(elem T) bool {
		return elem == match
	}

	return FindFn(fn, elems...)
}

func FindFn[T any](fn func(T) bool, elems ...T) (T, bool) {
	var ok bool
	var found T

	for _, elem := range elems {
		if ok = fn(elem); ok {
			found = elem
			break
		}
	}

	return found, ok
}

func FindFnErr[T any](fn func(T) (bool, error), elems ...T) (T, bool, error) {
	var found T
	var ok bool
	var err error

	for _, elem := range elems {
		if ok, err = fn(elem); ok || err != nil {
			found = elem
			break
		}
	}

	return found, ok, err
}
