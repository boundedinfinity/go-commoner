package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Dedup[T comparable](elems ...T) []T {
	return slicer.Dedup(elems...)
}

func DedupFn[T any, K comparable](fn func(T) K, elems ...T) []T {
	return slicer.DedupFn(fn, elems...)
}

func DedupFnErr[T any, K comparable](fn func(T) (K, error), elems ...T) trier.Try[[]T] {
	return trier.Complete(slicer.DedupFnErr(fn, elems...))
}
