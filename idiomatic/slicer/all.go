package slicer

// All test if the match value is equal to all elements in elems
func All[T comparable](match T, elems ...T) bool {
	fn := func(_ int, elem T) bool {
		return match == elem
	}

	return AllFn(fn, elems...)
}

func AllFn[T any](fn func(int, T) bool, elems ...T) bool {
	fn2 := func(i int, elem T) (bool, error) {
		return fn(i, elem), nil
	}

	all, _ := AllFnErr(fn2, elems...)
	return all
}

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
