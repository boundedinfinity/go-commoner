package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func FoldTry[I any, O any](initial O, vs []I, fn func(O, I) (O, error)) trier.Try[O] {
	return FoldLeftTry(initial, vs, fn)
}

func FoldLeftTry[I any, O any](initial O, vs []I, fn func(O, I) (O, error)) trier.Try[O] {
	curr, err := slicer.FoldLeftErr(initial, vs, fn)

	if err != nil {
		trier.Complete(curr, err)
	}

	return trier.Success(curr)
}

func FoldRightTry[I any, O any](initial O, vs []I, fn func(O, I) (O, error)) trier.Try[O] {
	curr, err := slicer.FoldRightErr(initial, vs, fn)
	return trier.Complete(curr, err)
}
