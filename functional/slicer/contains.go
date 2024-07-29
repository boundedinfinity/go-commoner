package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Contains[T comparable](v T, elems ...T) bool {
	return slicer.AnyOf(v, elems...)
}

func ContainsFn[T any](fn func(int, T) bool, elems ...T) bool {
	return slicer.AnyOfFn(fn, elems...)
}

func ContainsFnErr[T any](fn func(int, T) (bool, error), elems ...T) trier.Try[bool] {
	return trier.CompleteErr(slicer.AnyOfFnErr(fn, elems...))
}
