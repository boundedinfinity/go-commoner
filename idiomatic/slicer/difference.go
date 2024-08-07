package slicer

func Difference[T comparable](as, bs []T) []T {
	fn := func(_ int, t T) T { return t }
	return DifferenceFn(fn, as, bs)
}

func DifferenceFn[T any, C comparable](fn func(int, T) C, as, bs []T) []T {
	if fn == nil {
		return []T{}
	}

	fn2 := func(i int, t T) (C, error) {
		return fn(i, t), nil
	}

	results, _ := DifferenceFnErr(fn2, as, bs)
	return results
}

func DifferenceFnErr[T any, C comparable](fn func(int, T) (C, error), as, bs []T) ([]T, error) {
	results := []T{}
	var err error

	if fn == nil {
		return results, err
	}

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
