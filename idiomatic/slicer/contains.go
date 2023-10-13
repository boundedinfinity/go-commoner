package slicer

func Contains[T comparable](match T, items ...T) bool {
	fn := func(current bool, item T) bool {
		if current {
			return current
		}

		return match == item
	}

	return FoldLeft(false, fn, items...)
}

func ContainsFn[T comparable](fn func(T) bool, items ...T) bool {
	fn2 := func(current bool, item T) bool {
		if current {
			return current
		}

		return fn(item)
	}

	return FoldLeft(false, fn2, items...)
}

func ContainsFnI[T comparable](fn func(int, T) bool, items ...T) bool {
	fn2 := func(i int, current bool, item T) bool {
		if current {
			return current
		}

		return fn(i, item)
	}

	return FoldLeftI(false, fn2, items...)
}

func ContainsFnErr[T comparable](fn func(T) (bool, error), items ...T) (bool, error) {
	fn2 := func(current bool, item T) (bool, error) {
		if current {
			return current, nil
		}

		return fn(item)
	}

	return FoldLeftErr(false, fn2, items...)
}

func ContainsFnErrI[T comparable](fn func(int, T) (bool, error), items ...T) (bool, error) {
	fn2 := func(i int, current bool, item T) (bool, error) {
		if current {
			return current, nil
		}

		return fn(i, item)
	}

	return FoldLeftErrI(false, fn2, items...)
}
