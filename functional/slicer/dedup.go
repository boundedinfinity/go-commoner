package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Dedup[T comparable](items ...T) []T {
	return slicer.Dedup(items...)
}

func DedupFn[T any, K comparable](fn func(T) K, items ...T) []T {
	return slicer.DedupFn(fn, items...)
}

func DedupFnErr[T any, K comparable](fn func(T) (K, error), items ...T) trier.Try[[]T] {
	return trier.Complete(slicer.DedupFnErr(fn, items...))
}
