package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Fold[I any, O any](initial O, fn func(int, O, I) O, elems ...I) O {
	return slicer.Fold(fn, initial, elems...)
}

func FoldErr[I any, O any](initial O, fn func(int, O, I) (O, error), elems ...I) trier.Try[O] {
	return trier.Complete(slicer.FoldErr(fn, initial, elems...))
}

func FoldLeft[I any, O any](initial O, fn func(int, O, I) O, elems ...I) O {
	return slicer.FoldLeft(fn, initial, elems...)
}

func FoldLeftErr[I any, O any](initial O, fn func(int, O, I) (O, error), elems ...I) trier.Try[O] {
	return trier.Complete(slicer.FoldLeftErr(fn, initial, elems...))
}

func FoldRight[I any, O any](initial O, fn func(int, O, I) O, elems ...I) O {
	return slicer.FoldRight(fn, initial, elems...)
}

func FoldRightErr[I any, O any](initial O, fn func(int, O, I) (O, error), elems ...I) trier.Try[O] {
	return trier.Complete(slicer.FoldRightErr(fn, initial, elems...))
}
