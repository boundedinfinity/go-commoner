package slicer

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// References
//      - https://lodash.com/docs/4.17.15#includes
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// NoneOf[T] test if the match value is equal to any of the elements in elems
func NoneOf[T comparable](match T, elems ...T) bool {
	return !AnyOf(match, elems...)
}

// NoneOfFn[T] test if the value returned from fn is true for any of the elements in elems
//
// The fn(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
func NoneOfFn[T any](fn func(int, T) bool, elems ...T) bool {
	return !AnyOfFn(fn, elems...)
}

// NoneOfFnErr[T] test if the value returned from fn is true for any of the elements in elems
//
// The fn(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
//
// If the fn function returns an error, processing through elems is stopped and the error is returned.
func NoneOfFnErr[T any](fn func(int, T) (bool, error), elems ...T) (bool, error) {
	found, err := AnyOfFnErr(fn, elems...)
	return !found, err
}
