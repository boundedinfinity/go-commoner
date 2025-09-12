package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Dedup[T comparable](elems ...T) []T {
	return slicer.Deduplicate(elems...)
}

func DedupFn[T any, C comparable](fn func(T) C, elems ...T) []T {
	return slicer.DeduplicateBy(fn, elems...)
}

func DedupFnErr[T any, C comparable](fn func(T) (C, error), elems ...T) trier.Try[[]T] {
	return trier.CompleteErr(slicer.DeduplicateByErr(fn, elems...))
}
