package slicer

func Reject[T any](fn func(int, T) bool, elems ...T) []T {
	fn2 := func(i int, elem T) (bool, error) {
		return fn(i, elem), nil
	}

	results, _ := RejectErr(fn2, elems...)
	return results
}

func RejectErr[T any](fn func(int, T) (bool, error), elems ...T) ([]T, error) {
	var results []T

	for i, elem := range elems {
		ok, err := fn(i, elem)

		if err != nil {
			return results, err
		}

		if ok {
			results = append(results, elem)
		}
	}

	return results, nil
}
