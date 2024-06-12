package slicer

func SymmetricDifference[T comparable](as, bs []T) []T {
	fn := func(_ int, t T) T { return t }
	return SymmetricDifferenceFn(fn, as, bs)
}

func SymmetricDifferenceFn[T any, C comparable](fn func(int, T) C, as, bs []T) []T {
	fn2 := func(i int, t T) (C, error) {
		return fn(i, t), nil
	}

	results, _ := SymmetricDifferenceFnErr(fn2, as, bs)
	return results
}

func SymmetricDifferenceFnErr[T any, C comparable](fn func(int, T) (C, error), as, bs []T) ([]T, error) {
	results := []T{}

	xm, err := GroupErr(fn, append(as, bs...)...)
	if err != nil {
		return results, err
	}

	for _, xs := range xm {
		if len(xs) <= 1 {
			results = append(results, xs...)
		}
	}

	return results, nil
}
