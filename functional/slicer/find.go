package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Find[T comparable](match T, elems ...T) (T, bool) {
	return slicer.Find(match, elems...)
}

func FindFn[T any](fn func(T) bool, elems ...T) (T, bool) {
	return slicer.FindBy(fn, elems...)
}

func FindFnErr[T any](fn func(T) (bool, error), elems ...T) trier.Try[optioner.Option[T]] {
	t, ok, err := slicer.FindByErr(fn, elems...)
	return trier.CompleteErr(optioner.OfOk(t, ok), err)
}
