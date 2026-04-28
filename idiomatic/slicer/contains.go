package slicer

import "slices"

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// References
//      - https://lodash.com/docs/4.17.15#includes
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Contains test if the target value is equal to any of the elements in elems.
//
// Returns true on the first element that matches.
func Contains[T comparable](target T, elems ...T) bool {
	return slices.Contains(elems, target)
}

// ContainsFunc test if the value returned from fn is true for any of the elements in elems
//
// The by(T) bool functions takes:
//   - T is the current element
//
// and returns true if the the value of by(T) matches.
//
// Returns true on the first element that matches.
func ContainsFunc[T any](fn func(T) bool, elems ...T) bool {
	var ok bool

	for _, elem := range elems {
		if ok = fn(elem); ok {
			break
		}
	}

	return ok
}

// ContainsFuncI test if the value returned from fn is true for any of the elements in elems
//
// The by(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
//
// Returns true on the first element that matches.
func ContainsFuncI[T any](fn func(int, T) bool, elems ...T) bool {
	var ok bool

	for i, elem := range elems {
		if ok = fn(i, elem); ok {
			break
		}
	}

	return ok
}

// ContainsFuncErr test if the value returned from fn is true for any of the elements in elems
//
// The by(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
//
// Returns true on the first element that matches.
//
// If the fn function returns an error, processing through elems is stopped and the error is returned.
func ContainsFuncErr[T any](fn func(T) (bool, error), elems ...T) (bool, error) {
	var ok bool
	var err error

	for _, elem := range elems {
		if ok, err = fn(elem); ok || err != nil {
			break
		}
	}

	return ok, err
}

// ContainsFuncErrI test if the value returned from fn is true for any of the elements in elems
//
// The by(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
//
// Returns true on the first element that matches.
//
// If the fn function returns an error, processing through elems is stopped and the error is returned.
func ContainsFuncErrI[T any](fn func(int, T) (bool, error), elems ...T) (bool, error) {
	var ok bool
	var err error

	for i, elem := range elems {
		if ok, err = fn(i, elem); ok || err != nil {
			break
		}
	}

	return ok, err
}
