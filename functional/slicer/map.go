package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Map[T any, U any](fn func(T) U, items ...T) []U {
	return slicer.Map(fn, items...)
}

func MapI[T any, U any](fn func(int, T) U, items ...T) []U {
	return slicer.MapI(fn, items...)
}

func MapErr[T any, U any](fn func(T) (U, error), items ...T) trier.Try[[]U] {
	return trier.Complete(slicer.MapErr(fn, items...))
}

func MapErrI[T any, U any](fn func(int, T) (U, error), items ...T) trier.Try[[]U] {
	return trier.Complete(slicer.MapErrI(fn, items...))
}
