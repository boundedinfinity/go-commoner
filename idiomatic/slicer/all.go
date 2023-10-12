package slicer

func All[T comparable](x T, items ...T) bool {
	fn := func(v T) bool {
		return v == x
	}

	return AllFn(fn, items...)
}

func AllFn[T comparable](fn func(T) bool, items ...T) bool {
	for _, item := range items {
		if !fn(item) {
			return false
		}
	}

	return true
}

func AllFnI[T comparable](fn func(int, T) bool, items ...T) bool {
	for i, item := range items {
		if !fn(i, item) {
			return false
		}
	}

	return true
}

func AllFnErr[T comparable](fn func(T) (bool, error), items ...T) (bool, error) {
	ok := true
	var err error

	for _, item := range items {
		ok, err = fn(item)

		if !ok || err != nil {
			break
		}
	}

	return ok, err
}

func AllFnErrI[T comparable](fn func(int, T) (bool, error), items ...T) (bool, error) {
	ok := true
	var err error

	for i, item := range items {
		ok, err = fn(i, item)

		if !ok || err != nil {
			break
		}
	}

	return ok, err
}
