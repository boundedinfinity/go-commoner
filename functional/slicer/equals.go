package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Equals[T comparable](as, bs []T) bool {
	return slicer.Equals(as, bs)
}

func EqualsFn[T comparable](as, bs []T, fn func(T, T) bool) bool {
	return slicer.EqualsFn(as, bs, fn)
}

func EqualsFnErr[T comparable](as, bs []T, fn func(T, T) (bool, error)) trier.Try[bool] {
	return trier.Complete(slicer.EqualsFnErr(as, bs, fn))
}
