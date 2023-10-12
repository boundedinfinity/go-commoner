package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Diff[T comparable](as, bs []T) []T {
	return slicer.Diff(as, bs)
}

func DiffFn[T any](as, bs []T, fn func(T, T) bool) []T {
	return slicer.DiffFn(as, bs, fn)
}

func DiffFnErr[T any](as, bs []T, fn func(T, T) (bool, error)) trier.Try[[]T] {
	return trier.Complete(slicer.DiffFnErr(as, bs, fn))
}
