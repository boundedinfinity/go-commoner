package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Dedup[T comparable](elems ...T) []T {
	return slicer.Deduplicate(elems...)
}

func DedupFn[T any, K comparable](fn func(int, T) K, elems ...T) []T {
	return slicer.DeduplicateFn(fn, elems...)
}

func DedupFnErr[T any, K comparable](fn func(int, T) (K, error), elems ...T) trier.Try[[]T] {
	return trier.Complete(slicer.DeduplicateFnErr(fn, elems...))
}
