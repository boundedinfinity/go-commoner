package slicer

// All test if the target value is equal to all elements in elems
func All[T comparable](target T, elems ...T) bool {
	for _, elem := range elems {
		if elem != target {
			return false
		}
	}

	return true
}

// AllFunc test if the result of by(T) evaluates to true for all elements in elems
//
// The fn(T) bool functions takes:
//   - T is the current element in elems
func AllFunc[T any](fn func(T) bool, elems ...T) bool {
	if fn == nil {
		return false
	}

	for _, elem := range elems {
		if !fn(elem) {
			return false
		}
	}

	return true
}

// AllFuncI test if the result of fn(int, T) evaluates to true for all elements in elems
//
// The fn(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element in elems
func AllFuncI[T any](fn func(int, T) bool, elems ...T) bool {
	if fn == nil {
		return false
	}

	for i, elem := range elems {
		if fn(i, elem) {
			return true
		}
	}

	return false
}

// AllFuncErr test if the result of fn(T) err evaluates to true for all elements in elems
//
// The fn(T) bool functions takes:
//   - T is the current element
//
// If the fn function returns an error, processing through elems is stopped and the error is returned.
func AllFuncErr[T any](fn func(T) (bool, error), elems ...T) (bool, error) {
	if fn == nil {
		return false, nil
	}

	all := true
	var err error

	for _, elem := range elems {
		all, err = fn(elem)

		if err != nil {
			all = false
			break
		}

		if !all {
			break
		}
	}

	return all, err
}

// AllFuncErrI test if the result of fn(int, T) evaluates to true for all elements in elems
//
// The fn(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
//
// If the fn function returns an error, processing through elems is stopped and the error is returned.
func AllFuncErrI[T any](fn func(int, T) (bool, error), elems ...T) (bool, error) {
	if fn == nil {
		return false, nil
	}

	all := true
	var err error

	for i, elem := range elems {
		all, err = fn(i, elem)

		if err != nil {
			all = false
			break
		}

		if !all {
			break
		}
	}

	return all, err
}
