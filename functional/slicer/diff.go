package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Diff[T comparable](as, bs []T) optioner.Option[[]T] {
	return optioner.OfSlice(slicer.Diff(as, bs))
}

func DiffFn[T any](as, bs []T, fn func(T, T) bool) optioner.Option[[]T] {
	return optioner.OfSlice(slicer.DiffFn(as, bs, fn))
}

func DiffFnErr[T any](as, bs []T, fn func(T, T) (bool, error)) trier.Try[optioner.Option[[]T]] {
	return trier.CompleteOfSlice(slicer.DiffFnErr(as, bs, fn))
}
