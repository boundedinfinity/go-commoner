package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Exist[T comparable](s T, elems ...T) bool {
	return slicer.Exist(s, elems...)
}

func ExistFn[T any](fn func(T) bool, elems ...T) bool {
	return slicer.ExistBy(fn, elems...)
}

func ExistFnErr[T any](fn func(T) (bool, error), elems ...T) trier.Try[bool] {
	return trier.CompleteErr(slicer.ExistByErr(fn, elems...))
}
