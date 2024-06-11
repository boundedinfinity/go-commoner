package slicer

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// References
//      - https://lodash.com/docs/4.17.15#includes
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Contains test wheather the match value is equal to any of the elems.
func Contains[T comparable](match T, elems ...T) bool {
	fn := func(_ int, current bool, elem T) bool {
		if current {
			return current
		}

		return match == elem
	}

	return FoldLeft(false, fn, elems...)
}

func ContainsFn[T any](fn func(int, T) bool, elems ...T) bool {
	fn2 := func(i int, current bool, elem T) bool {
		if current {
			return current
		}

		return fn(i, elem)
	}

	return FoldLeft(false, fn2, elems...)
}

// ContainsFnErr test wheather the match value is equal to any of the elems.
//
// If the fn function returns an error, processing through elements is stopped and the error is returned.  The values
// of the grouped elements will contain any elements processed up to but not including the element when the error
// occured.
func ContainsFnErr[T any](fn func(int, T) (bool, error), elems ...T) (bool, error) {
	var found bool
	var err error

	for i, elem := range elems {
		found, err = fn(i, elem)

		if err != nil {
			break
		}

		if found {
			break
		}
	}

	return found, err
}
