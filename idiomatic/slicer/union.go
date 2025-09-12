package slicer

func Union[T comparable](as []T, bs []T) ([]T, error) {
	fn2 := func(_ int, elem T) T { return elem }
	return UnionFn(fn2, as, bs)
}

func UnionFn[T any, C comparable](fn func(int, T) C, as []T, bs []T) ([]T, error) {
	fn2 := func(i int, elem T) (C, error) {
		return fn(i, elem), nil
	}

	return UnionFnErr(fn2, as, bs)
}

func UnionFnErr[T any, C comparable](fn func(int, T) (C, error), as []T, bs []T) ([]T, error) {
	results := append([]T{}, as...)
	results = append(results, bs...)
	return UniqByErrI(fn, results...)
}
