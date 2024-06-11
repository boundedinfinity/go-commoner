package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Filter[T any](fn func(T) bool, elems ...T) []T {
	return slicer.Filter(fn, elems...)
}

func FilterErr[T any](fn func(T) (bool, error), elems ...T) trier.Try[[]T] {
	return trier.Complete(slicer.FilterErr(fn, elems...))
}
