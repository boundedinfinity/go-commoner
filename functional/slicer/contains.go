package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Contains[T comparable](v T, xs ...T) bool {
	return slicer.Contains(v, xs...)
}

func ContainsFn[T comparable](fn func(T) bool, xs ...T) bool {
	return slicer.ContainsFn(fn, xs...)
}

func ContainsFnErr[T comparable](fn func(T) (bool, error), xs ...T) trier.Try[bool] {
	return trier.Complete(slicer.ContainsFnErr(fn, xs...))
}
