package slicer

func Filter[T any](fn func(T) bool, elems ...T) []T {
	var os []T

	for _, elem := range elems {
		if fn(elem) {
			os = append(os, elem)
		}
	}

	return os
}

func FilterErr[T any](fn func(T) (bool, error), elems ...T) ([]T, error) {
	var os []T

	for _, elem := range elems {
		ok, err := fn(elem)

		if err != nil {
			return os, err
		}

		if ok {
			os = append(os, elem)
		}
	}

	return os, nil
}
