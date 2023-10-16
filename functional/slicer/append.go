package slicer

import "github.com/boundedinfinity/go-commoner/functional/optioner"

func Append[T any](slice optioner.Option[[]T], elems optioner.Option[[]T]) optioner.Option[[]T] {
	return optioner.OfSlice(append(slice.Get(), elems.Get()...))
}

func AppendLift[T any](slice optioner.Option[[]T], items ...T) optioner.Option[[]T] {
	return Append(slice, optioner.OfSlice(items))
}

func AppendOption[T any](slice optioner.Option[[]T], items ...optioner.Option[T]) optioner.Option[[]T] {
	return Append(slice, optioner.OfOptions(items...))
}
