package slicer

func All[T comparable](match T, elems ...T) bool {
	fn := func(_ int, current bool, elem T) bool {
		return current && match == elem
	}

	return FoldLeft(true, fn, elems...)
}

func AllFn[T any](fn func(int, T) bool, elems ...T) bool {
	fn2 := func(i int, current bool, elem T) bool {
		return current && fn(i, elem)
	}

	return FoldLeft(true, fn2, elems...)
}

func AllFnErr[T any](fn func(int, T) (bool, error), elems ...T) (bool, error) {
	fn2 := func(i int, current bool, elem T) (bool, error) {
		if ok, err := fn(i, elem); err != nil {
			return current, err
		} else {
			return current && ok, nil
		}
	}

	return FoldLeftErr(true, fn2, elems...)
}
