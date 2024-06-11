package slicer

func Filter[T any](fn func(int, T) bool, elems ...T) []T {
	fn2 := func(i int, elem T) (bool, error) {
		return fn(i, elem), nil
	}

	result, _ := FilterErr(fn2, elems...)
	return result
}

func FilterErr[T any](fn func(int, T) (bool, error), elems ...T) ([]T, error) {
	var os []T

	for i, elem := range elems {
		ok, err := fn(i, elem)

		if err != nil {
			return os, err
		}

		if ok {
			os = append(os, elem)
		}
	}

	return os, nil
}
