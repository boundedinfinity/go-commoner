package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Equals[T comparable](as, bs []T) bool {
	return slicer.Equals(as, bs)
}

func EqualsFn[T any, C comparable](fn func(int, T) C, as, bs []T) bool {
	return slicer.EqualsFn(fn, as, bs)
}

func EqualsFnErr[T comparable, C comparable](fn func(int, T) (C, error), as, bs []T) trier.Try[bool] {
	return trier.CompleteErr(slicer.EqualsFnErr(fn, as, bs))
}
