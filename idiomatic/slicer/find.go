package slicer

func FindFn[T any](fn func(T) bool, elems ...T) (T, bool) {
	for _, elem := range elems {
		if fn(elem) {
			return elem, true
		}
	}

	var zero T
	return zero, false
}

func FindFnI[T any](fn func(int, T) bool, elems ...T) (T, bool) {
	for i, elem := range elems {
		if fn(i, elem) {
			return elem, true
		}
	}

	var zero T
	return zero, false
}

func FindFnErr[T any](fn func(T) (bool, error), elems ...T) (T, bool, error) {
	var found T
	var ok bool
	var err error

	if fn == nil {
		return found, ok, err
	}

	for _, elem := range elems {
		if ok, err = fn(elem); ok || err != nil {
			found = elem
			break
		}
	}

	return found, ok, err
}

func FindFnErrI[T any](fn func(int, T) (bool, error), elems ...T) (T, bool, error) {
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
