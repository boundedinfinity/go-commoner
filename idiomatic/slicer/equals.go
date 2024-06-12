package slicer

func Equals[T comparable](as, bs []T) bool {
	fn := func(_ int, t T) T { return t }
	return EqualsFn(fn, as, bs)
}

func EqualsFn[T any, K comparable](fn func(int, T) K, as, bs []T) bool {
	if fn == nil {
		return false
	}

	fn2 := func(i int, elem T) (K, error) {
		return fn(i, elem), nil
	}

	result, _ := EqualsFnErr(fn2, as, bs)
	return result
}

func EqualsFnErr[T any, K comparable](fn func(int, T) (K, error), as, bs []T) (bool, error) {
	var ok bool
	var err error

	if fn == nil {
		return ok, err
	}

	return ok, err
}
