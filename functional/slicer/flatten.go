package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Flatten[T any](items ...[]T) []T {
	return slicer.Flatten(items...)
}

func FlattenFn[T any, U any](fn func(vs []T) []U, items ...[]T) []U {
	return slicer.FlattenFn(fn, items...)
}

func FlattenFnErr[T any, U any](fn func(vs []T) ([]U, error), items ...[]T) trier.Try[[]U] {
	return trier.Complete(slicer.FlattenFnErr(fn, items...))
}
