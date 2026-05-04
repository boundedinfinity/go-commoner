package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func All[T comparable](x T, vs ...T) bool {
	return slicer.All(x, vs...)
}

func AllBy[T any](fn func(T) bool, vs ...T) bool {
	return slicer.AllFunc(fn, vs...)
}

func AllByI[T any](fn func(int, T) bool, vs ...T) bool {
	return slicer.AllFuncI(fn, vs...)
}

func AllByErr[T any](fn func(T) (bool, error), vs ...T) trier.Try[bool] {
	return trier.Complete(slicer.AllFuncErr(fn, vs...))
}

func AllByErrI[T any](fn func(int, T) (bool, error), vs ...T) trier.Try[bool] {
	return trier.Complete(slicer.AllFuncErrI(fn, vs...))
}
