package slicer

func All[T comparable](match T, items ...T) bool {
	fn := func(current bool, item T) bool {
		return current && match == item
	}

	return FoldLeft(true, fn, items...)
}

func AllFn[T comparable](fn func(T) bool, items ...T) bool {
	fn2 := func(current bool, item T) bool {
		return current && fn(item)
	}

	return FoldLeft(true, fn2, items...)
}

func AllFnI[T comparable](fn func(int, T) bool, items ...T) bool {
	fn2 := func(i int, current bool, item T) bool {
		return current && fn(i, item)
	}

	return FoldLeftI(true, fn2, items...)
}

func AllFnErr[T comparable](fn func(T) (bool, error), items ...T) (bool, error) {
	fn2 := func(current bool, item T) (bool, error) {
		if ok, err := fn(item); err != nil {
			return current, err
		} else {
			return current && ok, nil
		}
	}

	return FoldLeftErr(true, fn2, items...)
}

func AllFnErrI[T comparable](fn func(int, T) (bool, error), items ...T) (bool, error) {
	fn2 := func(i int, current bool, item T) (bool, error) {
		if ok, err := fn(i, item); err != nil {
			return current, err
		} else {
			return current && ok, nil
		}
	}

	return FoldLeftErrI(true, fn2, items...)
}
