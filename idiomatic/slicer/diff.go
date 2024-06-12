package slicer

func Diff[T comparable](as, bs []T) []T {
	fn := func(_ int, t T) T { return t }
	return DiffFn(fn, as, bs)
}

func DiffFn[T any, C comparable](fn func(int, T) C, as, bs []T) []T {
	fn2 := func(i int, t T) (C, error) {
		return fn(i, t), nil
	}

	results, _ := DiffFnErr(fn2, as, bs)
	return results
}

func DiffFnErr[T any, C comparable](fn func(int, T) (C, error), as, bs []T) ([]T, error) {
	results := []T{}
	var err error

	for ai, a := range as {
		ac, ferr := fn(ai, a)

		if ferr != nil {
			err = ferr
			break
		}

		var found bool

		for _, b := range bs {
			bc, ferr := fn(ai, b)

			if ferr != nil {
				err = ferr
				break
			}

			if ac == bc {
				found = true
			}
		}

		if err != nil {
			break
		}

		if !found {
			results = append(results, a)
		}
	}

	return results, err
}
