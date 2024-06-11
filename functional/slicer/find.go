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
	return slicer.FindFn(fn, elems...)
}

func FindFnErr[T any](fn func(T) (bool, error), elems ...T) trier.Try[optioner.Option[T]] {
	t, ok, err := slicer.FindFnErr(fn, elems...)
	return trier.Complete(optioner.OfOk(t, ok), err)
}
