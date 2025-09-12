package slicer

// All test if the target value is equal to all elements in elems
func All[T comparable](target T, elems ...T) bool {
	for _, elem := range elems {
		if elem == target {
			return true
		}
	}

	return false
}

// AllBy test if the result of by(T) evaluates to true for all elements in elems
//
// The fn(T) bool functions takes:
//   - T is the current element in elems
func AllBy[T any](by func(T) bool, elems ...T) bool {
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

// AllByI test if the result of fn(int, T) evaluates to true for all elements in elems
//
// The by(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element in elems
func AllByI[T any](by func(int, T) bool, elems ...T) bool {
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

// AllByErr test if the result of by(T) err evaluates to true for all elements in elems
//
// The by(int, T) bool functions takes:
//   - T is the current element
//
// If the by function returns an error, processing through elems is stopped and the error is returned.
func AllByErr[T any](by func(T) (bool, error), elems ...T) (bool, error) {
	if by == nil {
		return false, nil
	}

	all := true
	var err error

	for _, elem := range elems {
		all, err = by(elem)

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

// AllByErrI test if the result of fn(int, T) evaluates to true for all elements in elems
//
// The fn(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
//
// If the fn function returns an error, processing through elems is stopped and the error is returned.
func AllByErrI[T any](by func(int, T) (bool, error), elems ...T) (bool, error) {
	if by == nil {
		return false, nil
	}

	all := true
	var err error

	for i, elem := range elems {
		all, err = by(i, elem)

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
