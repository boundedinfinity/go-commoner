package slicer

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

func Chunk[T comparable](size int, item ...T) [][]T {
	return slicer.Chunk(size, item...)
}
