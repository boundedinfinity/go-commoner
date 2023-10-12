package slicer

func Contains[T comparable](v T, items ...T) bool {
	fn := func(x T) bool {
		return x == v
	}

	return ContainsFn(fn, items...)
}

func ContainsFn[T comparable](fn func(T) bool, items ...T) bool {
	for _, x := range items {
		if fn(x) {
			return true
		}
	}

	return false
}

func ContainsFnI[T comparable](fn func(int, T) bool, items ...T) bool {
	for i, x := range items {
		if fn(i, x) {
			return true
		}
	}

	return false
}

func ContainsFnErr[T comparable](fn func(T) (bool, error), items ...T) (bool, error) {
	var ok bool
	var err error

	for _, x := range items {
		if ok, err = fn(x); ok || err != nil {
			break
		}
	}

	return ok, err
}

func ContainsFnErrI[T comparable](fn func(int, T) (bool, error), items ...T) (bool, error) {
	var ok bool
	var err error

	for i, x := range items {
		if ok, err = fn(i, x); ok || err != nil {
			break
		}
	}

	return ok, err
}
