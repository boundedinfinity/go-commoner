package slicer

func Exist[T comparable](s T, items ...T) bool {
	fn := func(_ int, t T) bool {
		return t == s
	}

	return ExistFn(fn, items...)
}

func ExistFn[T any](fn func(int, T) bool, items ...T) bool {
	wrap := func(i int, t T) (bool, error) {
		return fn(i, t), nil
	}

	found, _ := ExistFnErr(wrap, items...)
	return found
}

func ExistFnErr[T any](fn func(int, T) (bool, error), items ...T) (bool, error) {
	var ok bool
	var err error

	for i, item := range items {
		ok, err = fn(i, item)

		if !ok {
			break
		}

		if err != nil {
			break
		}
	}

	return ok, err
}
