package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func FoldLeft[I any, O any](initial O, fn func(O, I) O, items ...I) O {
	return slicer.FoldLeft(initial, fn, items...)
}

func FoldLeftI[I any, O any](initial O, fn func(int, O, I) O, items ...I) O {
	return slicer.FoldLeftI(initial, fn, items...)
}

func FoldLeftErr[I any, O any](initial O, fn func(O, I) (O, error), items ...I) trier.Try[O] {
	return trier.Complete(slicer.FoldLeftErr(initial, fn, items...))
}

func FoldLeftErrI[I any, O any](initial O, fn func(int, O, I) (O, error), items ...I) trier.Try[O] {
	return trier.Complete(slicer.FoldLeftErrI(initial, fn, items...))
}

func FoldRight[I any, O any](initial O, fn func(O, I) O, items ...I) O {
	return slicer.FoldRight(initial, fn, items...)
}

func FoldRightI[I any, O any](initial O, fn func(int, O, I) O, items ...I) O {
	return slicer.FoldRightI(initial, fn, items...)
}

func FoldRightErr[I any, O any](initial O, fn func(O, I) (O, error), items ...I) trier.Try[O] {
	return trier.Complete(slicer.FoldRightErr(initial, fn, items...))
}

func FoldRightErrI[I any, O any](initial O, fn func(int, O, I) (O, error), items ...I) trier.Try[O] {
	return trier.Complete(slicer.FoldRightErrI(initial, fn, items...))
}
