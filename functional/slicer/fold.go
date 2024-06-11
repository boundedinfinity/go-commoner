package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Fold[I any, O any](initial O, fn func(int, O, I) O, elems ...I) O {
	return slicer.Fold(initial, fn, elems...)
}

func FoldErr[I any, O any](initial O, fn func(int, O, I) (O, error), elems ...I) trier.Try[O] {
	return trier.Complete(slicer.FoldErr(initial, fn, elems...))
}

func FoldLeft[I any, O any](initial O, fn func(int, O, I) O, elems ...I) O {
	return slicer.FoldLeft(initial, fn, elems...)
}

func FoldLeftErr[I any, O any](initial O, fn func(int, O, I) (O, error), elems ...I) trier.Try[O] {
	return trier.Complete(slicer.FoldLeftErr(initial, fn, elems...))
}

func FoldRight[I any, O any](initial O, fn func(int, O, I) O, elems ...I) O {
	return slicer.FoldRight(initial, fn, elems...)
}

func FoldRightErr[I any, O any](initial O, fn func(int, O, I) (O, error), elems ...I) trier.Try[O] {
	return trier.Complete(slicer.FoldRightErr(initial, fn, elems...))
}
