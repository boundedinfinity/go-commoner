package slicer

// All[T] test if the match value is equal to all elements in elems
func All[T comparable](match T, elems ...T) bool {
	fn := func(_ int, elem T) bool { return match == elem }
	return AllFn(fn, elems...)
}

// All[T] test if the result of fn(int, T) evaluates to true for all elements in elems
//
// The fn(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
func AllFn[T any](fn func(int, T) bool, elems ...T) bool {
	fn2 := func(i int, elem T) (bool, error) {
		return fn(i, elem), nil
	}

	all, _ := AllFnErr(fn2, elems...)
	return all
}

// All[T] test if the result of fn(int, T) evaluates to true for all elements in elems
//
// The fn(int, T) bool functions takes:
//   - int is the index of the current element
//   - T is the current element
//
// If the fn function returns an error, processing through elems is stopped and the error is returned.
func AllFnErr[T any](fn func(int, T) (bool, error), elems ...T) (bool, error) {
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
