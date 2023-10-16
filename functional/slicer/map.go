package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Map[T any, U any](fn func(T) U, items optioner.Option[[]T]) optioner.Option[[]U] {
	if items.Defined() {
		return optioner.OfSlice(slicer.Map(fn, items.Get()...))
	}

	return optioner.None[[]U]()
}

func MapI[T any, U any](fn func(int, T) U, items optioner.Option[[]T]) optioner.Option[[]U] {
	if items.Defined() {
		return optioner.OfSlice(slicer.MapI(fn, items.Get()...))
	}

	return optioner.None[[]U]()
}

func MapErr[T any, U any](fn func(T) (U, error), items optioner.Option[[]T]) trier.Try[optioner.Option[[]U]] {
	if items.Defined() {
		results, err := slicer.MapErr(fn, items.Get()...)

		return trier.Complete(optioner.OfSlice(results), err)
	}

	return trier.Complete(optioner.None[[]U](), nil)
}

func MapErrI[T any, U any](fn func(int, T) (U, error), items optioner.Option[[]T]) trier.Try[optioner.Option[[]U]] {
	if items.Defined() {
		results, err := slicer.MapErrI(fn, items.Get()...)

		return trier.Complete(optioner.OfSlice(results), err)
	}

	return trier.Complete(optioner.None[[]U](), nil)
}
