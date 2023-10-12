package slicer

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

func Chunk[T comparable](xs []T, size int) [][]T {
	return slicer.Chunk(xs, size)
}
