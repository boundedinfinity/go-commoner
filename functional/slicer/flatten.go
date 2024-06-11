package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Flatten[T any](elems ...[]T) []T {
	return slicer.Flatten(elems...)
}

func FlattenFn[T any, U any](fn func(vs []T) []U, elems ...[]T) []U {
	return slicer.FlattenFn(fn, elems...)
}

func FlattenFnErr[T any, U any](fn func(vs []T) ([]U, error), elems ...[]T) trier.Try[[]U] {
	return trier.Complete(slicer.FlattenFnErr(fn, elems...))
}
