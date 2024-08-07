package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func All[T comparable](x T, vs ...T) bool {
	return slicer.All(x, vs...)
}

func AllFn[T any](fn func(int, T) bool, vs ...T) bool {
	return slicer.AllFn(fn, vs...)
}

func AllFnErr[T any](fn func(int, T) (bool, error), vs ...T) trier.Try[bool] {
	return trier.CompleteErr(slicer.AllFnErr(fn, vs...))
}
