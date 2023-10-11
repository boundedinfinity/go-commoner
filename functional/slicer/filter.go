package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func FilterTry[T any](xs []T, fn slicer.FilterErrFn[T]) trier.Try[[]T] {
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
