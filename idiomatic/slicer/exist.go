package slicer

func Exist[T comparable](s T, elems ...T) bool {
	fn := func(_ int, t T) bool {
		return t == s
	}

	return ExistFn(fn, elems...)
}

func ExistFn[T any](fn func(int, T) bool, elems ...T) bool {
	wrap := func(i int, t T) (bool, error) {
		return fn(i, t), nil
	}

	found, _ := ExistFnErr(wrap, elems...)
	return found
}

func ExistFnErr[T any](fn func(int, T) (bool, error), elems ...T) (bool, error) {
	var ok bool
	var err error

	for i, elem := range elems {
		ok, err = fn(i, elem)

		if !ok {
			break
		}

		if err != nil {
			break
		}
	}

	return ok, err
}
