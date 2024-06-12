package slicer

func Uniq[C comparable](elems ...C) []C {
	fn2 := func(_ int, elem C) C { return elem }
	return UniqFn(fn2, elems...)
}

func UniqFn[T any, C comparable](fn func(int, T) C, elems ...T) []T {
	fn2 := func(i int, elem T) (C, error) {
		return fn(i, elem), nil
	}

	results, _ := UniqFnErr(fn2, elems...)

	return results
}

func UniqFnErr[T any, C comparable](fn func(int, T) (C, error), elems ...T) ([]T, error) {
	results := []T{}
	m := map[C]T{}

	for i, elem := range elems {
		c, err := fn(i, elem)

		if err != nil {
			return results, err
		}

		if _, ok := m[c]; !ok {
			m[c] = elem
		}
	}

	for _, v := range m {
		results = append(results, v)
	}

	return results, nil
}
