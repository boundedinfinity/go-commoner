package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Contains[T comparable](v T, items ...T) bool {
	return slicer.Contains(v, items...)
}

func ContainsFn[T comparable](fn func(T) bool, items ...T) bool {
	return slicer.ContainsFn(fn, items...)
}

func ContainsFnI[T comparable](fn func(int, T) bool, items ...T) bool {
	return slicer.ContainsFnI(fn, items...)
}

func ContainsFnErr[T comparable](fn func(T) (bool, error), items ...T) trier.Try[bool] {
	return trier.Complete(slicer.ContainsFnErr(fn, items...))
}

func ContainsFnErrI[T comparable](fn func(int, T) (bool, error), items ...T) trier.Try[bool] {
	return trier.Complete(slicer.ContainsFnErrI(fn, items...))
}
