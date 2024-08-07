package slicer

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// References
//      - https://lodash.com/docs/4.17.15#includes
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// AnyOf[T] test if the match value is equal to any of the elements in elems
func AnyOf[T comparable](match T, elems ...T) bool {
	fn := func(_ int, elem T) bool {
		return match == elem
	}

	return AnyOfFn(fn, elems...)
}

// AnyOfFn[T] test if the value returned from fn is true for any of the elements in elems
//
// The fn(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
func AnyOfFn[T any](fn func(int, T) bool, elems ...T) bool {
	if fn == nil {
		return false
	}

	fn2 := func(i int, elem T) (bool, error) {
		return fn(i, elem), nil
	}

	found, _ := AnyOfFnErr(fn2, elems...)
	return found
}

// AnyOfFnErr[T] test if the value returned from fn is true for any of the elements in elems
//
// The fn(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
//
// If the fn function returns an error, processing through elems is stopped and the error is returned.
func AnyOfFnErr[T any](fn func(int, T) (bool, error), elems ...T) (bool, error) {
	var found bool
	var err error

	if fn == nil {
		return found, err
	}

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
