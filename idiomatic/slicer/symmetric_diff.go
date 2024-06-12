package slicer

func SymmetricDiff[T comparable](as, bs []T) []T {
	fn := func(_ int, t T) T { return t }
	return SymmetricDiffFn(fn, as, bs)
}

func SymmetricDiffFn[T any, C comparable](fn func(int, T) C, as, bs []T) []T {
	fn2 := func(i int, t T) (C, error) {
		return fn(i, t), nil
	}

	results, _ := SymmetricDiffFnErr(fn2, as, bs)
	return results
}

func SymmetricDiffFnErr[T any, C comparable](fn func(int, T) (C, error), as, bs []T) ([]T, error) {
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
