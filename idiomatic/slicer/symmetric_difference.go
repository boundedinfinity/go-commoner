package slicer

func SymmetricDifference[T comparable](as, bs []T) []T {
	count := Count(append(as, bs...)...)
	var found []T

	for _, r := range count {
		if r.Count == 1 {
			found = append(found, r.Item)
		}
	}

	return found
}

func SymmetricDifferenceFn[T any, C comparable](fn func(T) C, as, bs []T) []T {
	count := CountFn(fn, append(as, bs...)...)
	var found []T

	for _, r := range count {
		if r.Count == 1 {
			found = append(found, r.Item)
		}
	}

	return found
}

func SymmetricDifferenceFnI[T any, C comparable](fn func(int, T) C, as, bs []T) []T {
	count := CountFnI(fn, append(as, bs...)...)
	var found []T

	for _, r := range count {
		if r.Count == 1 {
			found = append(found, r.Item)
		}
	}

	return found
}

func SymmetricDifferenceFnErr[T any, C comparable](fn func(T) (C, error), as, bs []T) ([]T, error) {
	var found []T
	count, err := CountFnErr(fn, append(as, bs...)...)

	if err != nil {
		return found, err
	}

	for _, r := range count {
		if r.Count == 1 {
			found = append(found, r.Item)
		}
	}

	return found, nil
}

func SymmetricDifferenceFnErrI[T any, C comparable](fn func(int, T) (C, error), as, bs []T) ([]T, error) {
	var found []T
	count, err := CountFnErrI(fn, append(as, bs...)...)

	if err != nil {
		return found, err
	}

	for _, r := range count {
		if r.Count == 1 {
			found = append(found, r.Item)
		}
	}

	return found, nil
}
