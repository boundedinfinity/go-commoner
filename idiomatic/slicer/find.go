package slicer

func Find[T comparable](target T, elems ...T) (T, bool) {
	for _, elem := range elems {
		if target == elem {
			return target, true
		}
	}

	var zero T
	return zero, false
}

func FindBy[T any](by func(T) bool, elems ...T) (T, bool) {
	for _, elem := range elems {
		if by(elem) {
			return elem, true
		}
	}

	var zero T
	return zero, false
}

func FindByI[T any](by func(int, T) bool, elems ...T) (T, bool) {
	for i, elem := range elems {
		if by(i, elem) {
			return elem, true
		}
	}

	var zero T
	return zero, false
}

func FindByErr[T any](by func(T) (bool, error), elems ...T) (T, bool, error) {
	var found T
	var ok bool
	var err error

	if by == nil {
		return found, ok, err
	}

	for _, elem := range elems {
		if ok, err = by(elem); ok || err != nil {
			found = elem
			break
		}
	}

	return found, ok, err
}

func FindByErrI[T any](fn func(int, T) (bool, error), elems ...T) (T, bool, error) {
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
