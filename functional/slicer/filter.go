package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Filter[T any](fn func(T) bool, items ...T) []T {
	return slicer.Filter(fn, items...)
}

func FilterErr[T any](fn func(T) (bool, error), items ...T) trier.Try[[]T] {
	return trier.Complete(slicer.FilterErr(fn, items...))
}
