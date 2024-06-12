package slicer

func Filter[T any](fn func(int, T) bool, elems ...T) []T {
	if fn == nil {
		return []T{}
	}

	fn2 := func(i int, elem T) (bool, error) {
		return fn(i, elem), nil
	}

	result, _ := FilterErr(fn2, elems...)
	return result
}

func FilterErr[T any](fn func(int, T) (bool, error), elems ...T) ([]T, error) {
	var result []T

	if fn == nil {
		return result, nil
	}

	for i, elem := range elems {
		ok, err := fn(i, elem)

		if err != nil {
			return result, err
		}

		if ok {
			result = append(result, elem)
		}
	}

	return result, nil
}
