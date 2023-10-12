package slicer

func Filter[T any](fn func(T) bool, items ...T) []T {
	var os []T

	for _, item := range items {
		if fn(item) {
			os = append(os, item)
		}
	}

	return os
}

func FilterErr[T any](fn func(T) (bool, error), items ...T) ([]T, error) {
	var os []T

	for _, item := range items {
		ok, err := fn(item)

		if err != nil {
			return os, err
		}

		if ok {
			os = append(os, item)
		}
	}

	return os, nil
}
