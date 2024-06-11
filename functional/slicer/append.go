package slicer

import "github.com/boundedinfinity/go-commoner/functional/optioner"

func Append[T any](slice optioner.Option[[]T], elems optioner.Option[[]T]) optioner.Option[[]T] {
	return optioner.OfSlice(append(slice.Get(), elems.Get()...))
}

func AppendLift[T any](slice optioner.Option[[]T], elems ...T) optioner.Option[[]T] {
	return Append(slice, optioner.OfSlice(elems))
}

func AppendOption[T any](slice optioner.Option[[]T], elems ...optioner.Option[T]) optioner.Option[[]T] {
	return Append(slice, optioner.OfOptions(elems...))
}
