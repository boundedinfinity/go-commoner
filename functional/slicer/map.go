package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Map[T any, U any](fn func(int, T) U, elems optioner.Option[[]T]) optioner.Option[[]U] {
	if elems.Defined() {
		return optioner.OfSlice(slicer.Map(fn, elems.Get()...))
	}

	return optioner.None[[]U]()
}

func MapErr[T any, U any](fn func(int, T) (U, error), elems optioner.Option[[]T]) trier.Try[optioner.Option[[]U]] {
	if elems.Defined() {
		results, err := slicer.MapErr(fn, elems.Get()...)

		return trier.CompleteErr(optioner.OfSlice(results), err)
	}

	return trier.CompleteErr(optioner.None[[]U](), nil)
}
