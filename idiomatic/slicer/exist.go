package slicer

func Exist[T comparable](target T, elems ...T) bool {
	for _, elem := range elems {
		if elem == target {
			return true
		}
	}

	return false
}

func ExistBy[T any](by func(T) bool, elems ...T) bool {
	if by == nil {
		return false
	}

	for _, elem := range elems {
		if by(elem) {
			return true
		}
	}

	return false
}

func ExistByI[T any](by func(int, T) bool, elems ...T) bool {
	if by == nil {
		return false
	}

	for i, elem := range elems {
		if by(i, elem) {
			return true
		}
	}

	return false
}

func ExistByErr[T any](by func(T) (bool, error), elems ...T) (bool, error) {
	var ok bool
	var err error

	if by == nil {
		return ok, err
	}

	for _, elem := range elems {
		ok, err = by(elem)

		if ok {
			break
		}

		if err != nil {
			break
		}
	}

	return ok, err
}

func ExistByErrI[T any](by func(int, T) (bool, error), elems ...T) (bool, error) {
	var ok bool
	var err error

	if by == nil {
		return ok, err
	}

	for i, elem := range elems {
		ok, err = by(i, elem)

		if ok {
			break
		}

		if err != nil {
			break
		}
	}

	return ok, err
}
