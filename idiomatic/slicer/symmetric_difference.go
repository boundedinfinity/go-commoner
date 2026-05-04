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
	count := CountFunc(fn, append(as, bs...)...)
	var found []T

	for _, r := range count {
		if r.Count == 1 {
			found = append(found, r.Item)
		}
	}

	return found
}

func SymmetricDifferenceFuncI[T any, C comparable](fn func(int, T) C, as, bs []T) []T {
	count := CountFuncI(fn, append(as, bs...)...)
	var found []T

	for _, r := range count {
		if r.Count == 1 {
			found = append(found, r.Item)
		}
	}

	return found
}

func SymmetricDifferenceFuncErr[T any, C comparable](fn func(T) (C, error), as, bs []T) ([]T, error) {
	var found []T
	count, err := CountFuncErr(fn, append(as, bs...)...)

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

func SymmetricDifferenceFuncErrI[T any, C comparable](fn func(int, T) (C, error), as, bs []T) ([]T, error) {
	var found []T
	count, err := CountFuncErrI(fn, append(as, bs...)...)

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
