package slicer

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// References
//      - https://lodash.com/docs/4.17.15#includes
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// AnyOf test if the target value is equal to any of the elements in elems.
//
// Returns true on the first element that matches.
func AnyOf[T comparable](target T, elems ...T) bool {
	var ok bool

	for _, elem := range elems {
		if target == elem {
			ok = true
			break
		}
	}

	return ok
}

// AnyOfBy test if the value returned from fn is true for any of the elements in elems
//
// The by(T) bool functions takes:
//   - T is the current element
//
// and returns true if the the value of by(T) matches.
//
// Returns true on the first element that matches.
func AnyOfBy[T any](by func(T) bool, elems ...T) bool {
	var ok bool

	for _, elem := range elems {
		if ok = by(elem); ok {
			break
		}
	}

	return ok
}

// AnyOfByI test if the value returned from fn is true for any of the elements in elems
//
// The by(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
//
// Returns true on the first element that matches.
func AnyOfByI[T any](by func(int, T) bool, elems ...T) bool {
	var ok bool

	for i, elem := range elems {
		if ok = by(i, elem); ok {
			break
		}
	}

	return ok
}

// AnyOfByErr test if the value returned from fn is true for any of the elements in elems
//
// The by(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
//
// Returns true on the first element that matches.
//
// If the fn function returns an error, processing through elems is stopped and the error is returned.
func AnyOfByErr[T any](by func(T) (bool, error), elems ...T) (bool, error) {
	var ok bool
	var err error

	for _, elem := range elems {
		if ok, err = by(elem); ok || err != nil {
			break
		}
	}

	return ok, err
}

// AnyOfByErrI test if the value returned from fn is true for any of the elements in elems
//
// The by(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
//
// Returns true on the first element that matches.
//
// If the fn function returns an error, processing through elems is stopped and the error is returned.
func AnyOfByErrI[T any](by func(int, T) (bool, error), elems ...T) (bool, error) {
	var ok bool
	var err error

	for i, elem := range elems {
		if ok, err = by(i, elem); ok || err != nil {
			break
		}
	}

	return ok, err
}
