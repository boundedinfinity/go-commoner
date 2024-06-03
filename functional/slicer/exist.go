package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Exist[T comparable](s T, items ...T) bool {
	return slicer.Exist(s, items...)
}

func ExistFn[T any](fn func(int, T) bool, items ...T) bool {
	return slicer.ExistFn(fn, items...)
}

func ExistFnErr[T any](fn func(int, T) (bool, error), items ...T) trier.Try[bool] {
	return trier.Complete(slicer.ExistFnErr(fn, items...))
}
