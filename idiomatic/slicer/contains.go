package slicer

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// References
//      - https://lodash.com/docs/4.17.15#includes
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Contains[T] test if the match value is equal to any of the elements in elems
func Contains[T comparable](match T, elems ...T) bool {
	fn := func(_ int, current bool, elem T) bool {
		return match == elem
	}

	return FoldLeft(false, fn, elems...)
}

// Contains[T] test if the value returned from fn is true for any of the elements in elems
//
// The fn(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
func ContainsFn[T any](fn func(int, T) bool, elems ...T) bool {
	fn2 := func(i int, elem T) (bool, error) {
		return fn(i, elem), nil
	}

	found, _ := ContainsFnErr(fn2, elems...)
	return found
}

// Contains[T] test if the value returned from fn is true for any of the elements in elems
//
// The fn(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
//
// If the fn function returns an error, processing through elems is stopped and the error is returned.
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
