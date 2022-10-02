package slices

import "github.com/boundedinfinity/go-commoner/trier"

type FilterFn[T any] func(T) bool

func Filter[T any](xs []T, fn FilterFn[T]) []T {
	var os []T

	for _, x := range xs {
		if fn(x) {
			os = append(os, x)
		}
	}

	return os
}

type FilterErrFn[T any] func(T) (bool, error)

func FilterErr[T any](xs []T, fn FilterErrFn[T]) ([]T, error) {
	var os []T

	for _, x := range xs {
		ok, err := fn(x)

		if err != nil {
			return os, err
		}

		if ok {
			os = append(os, x)
		}
	}

	return os, nil
}

func FilterTry[T any](xs []T, fn FilterErrFn[T]) trier.Try[[]T] {
	var os []T

	for _, x := range xs {
		ok, err := fn(x)

		if err != nil {
			return trier.Complete(os, err)
		}

		if ok {
			os = append(os, x)
		}
	}

	return trier.Success(os)
}
