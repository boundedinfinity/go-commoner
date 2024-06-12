package slicer

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Flatten[T any](elems ...[]T) []T {
	return slicer.Flatten(elems...)
}
