package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Diff[T comparable](as, bs []T) optioner.Option[[]T] {
	return optioner.OfSlice(slicer.Difference(as, bs))
}

func DiffFn[T any](as, bs []T, fn func(int, T) bool) optioner.Option[[]T] {
	return optioner.OfSlice(slicer.DifferenceFn(fn, as, bs))
}

func DiffFnErr[T any](as, bs []T, fn func(int, T) (bool, error)) trier.Try[optioner.Option[[]T]] {
	return trier.CompleteSlice(slicer.DifferenceFnErr(fn, as, bs))
}
