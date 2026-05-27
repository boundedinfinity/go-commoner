package slicer

func Reject[T any](fn func(T) bool, elems ...T) []T {
	var results []T

	for _, elem := range elems {
		if !fn(elem) {
			results = append(results, elem)
		}
	}

	return results
}

func RejectI[T any](fn func(int, T) bool, elems ...T) []T {
	var results []T

	for i, elem := range elems {
		if !fn(i, elem) {
			results = append(results, elem)
		}
	}

	return results
}

func RejectErr[T any](fn func(T) (bool, error), elems ...T) ([]T, error) {
	var results []T

	for _, elem := range elems {
		ok, err := fn(elem)

		if err != nil {
			return results, err
		}

		if !ok {
			results = append(results, elem)
		}
	}

	return results, nil
}

func RejectIErr[T any](fn func(int, T) (bool, error), elems ...T) ([]T, error) {
	var results []T

	for i, elem := range elems {
		ok, err := fn(i, elem)

		if err != nil {
			return results, err
		}

		if !ok {
			results = append(results, elem)
		}
	}

	return results, nil
}
