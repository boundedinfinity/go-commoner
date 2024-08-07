package slicer

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Chunk[T any](size int, elem ...T) optioner.Option[[][]T] {
	return optioner.OfSlice(slicer.Chunk(size, elem...))
}
