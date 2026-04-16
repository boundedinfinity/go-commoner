package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Contains[T comparable](v T, elems ...T) bool {
	return slicer.Contains(v, elems...)
}

func ContainsBy[T any](fn func(T) bool, elems ...T) bool {
	return slicer.ContainsFn(fn, elems...)
}

func ContainsByI[T any](fn func(int, T) bool, elems ...T) bool {
	return slicer.CcontainsFnI(fn, elems...)
}

func ContainsByErr[T any](fn func(T) (bool, error), elems ...T) trier.Try[bool] {
	return trier.Complete(slicer.ContainsFnErr(fn, elems...))
}

func ContainsByErrI[T any](fn func(int, T) (bool, error), elems ...T) trier.Try[bool] {
	return trier.Complete(slicer.ContainsFnErrI(fn, elems...))
}
