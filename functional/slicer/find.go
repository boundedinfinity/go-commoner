package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Find[T comparable](match T, items ...T) (T, bool) {
	return slicer.Find(match, items...)
}

func FindFn[T any](fn func(T) bool, items ...T) (T, bool) {
	return slicer.FindFn(fn, items...)
}

func FindFnErr[T any](fn func(T) (bool, error), items ...T) trier.Try[optioner.Option[T]] {
	t, ok, err := slicer.FindFnErr(fn, items...)
	return trier.Complete(optioner.OfI(t, ok), err)
}
