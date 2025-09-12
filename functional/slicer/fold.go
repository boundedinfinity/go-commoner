package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func FoldI[I any, O any](initial O, fn func(int, O, I) O, elems ...I) O {
	return slicer.FoldI(fn, initial, elems...)
}

func FoldErrI[I any, O any](initial O, fn func(int, O, I) (O, error), elems ...I) trier.Try[O] {
	return trier.CompleteErr(slicer.FoldErrI(fn, initial, elems...))
}

func FoldRight[I any, O any](initial O, fn func(int, O, I) O, elems ...I) O {
	return slicer.FoldRightI(fn, initial, elems...)
}

func FoldRightErrI[I any, O any](initial O, fn func(int, O, I) (O, error), elems ...I) trier.Try[O] {
	return trier.CompleteErr(slicer.FoldRightErrI(fn, initial, elems...))
}
