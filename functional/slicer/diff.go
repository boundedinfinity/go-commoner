package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Diff[T comparable](as, bs []T) optioner.Option[[]T] {
	return optioner.OfSlice(slicer.Difference(as, bs))
}

func DiffBy[T any](as, bs []T, fn func(T) bool) optioner.Option[[]T] {
	return optioner.OfSlice(slicer.DifferenceBy(fn, as, bs))
}

func DiffByErr[T any](as, bs []T, fn func(T) (bool, error)) trier.Try[optioner.Option[[]T]] {
	return trier.CompleteSlice(slicer.DifferenceByErr(fn, as, bs))
}
