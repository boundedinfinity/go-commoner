package slicer

func Contains[T comparable](v T, xs ...T) bool {
	fn := func(x T) bool {
		return x == v
	}

	return ContainsFn(fn, xs...)
}

func ContainsFn[T comparable](fn func(T) bool, xs ...T) bool {
	for _, x := range xs {
		if fn(x) {
			return true
		}
	}

	return false
}

func ContainsFnErr[T comparable](fn func(T) (bool, error), xs ...T) (bool, error) {
	var ok bool
	var err error

	for _, x := range xs {
		ok, err = fn(x)

		if ok || err != nil {
			break
		}
	}

	return ok, err
}
