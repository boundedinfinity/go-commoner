package matrix

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/math"

	"golang.org/x/exp/constraints"
)

func SliceToDescriptor[T constraints.Integer, V any](slice [][]V) MatrixDescriptor[T] {
	descriptor := MatrixDescriptor[T]{
		Rows: T(len(slice)),
	}

	for _, row := range slice {
		descriptor.Cols = math.Max(descriptor.Cols, T(len(row)))
	}

	return descriptor
}

type MatrixDescriptor[T constraints.Integer] struct {
	Rows T
	Cols T
}

func (t MatrixDescriptor[T]) OrdinalToPosition(ord T) MatrixCoordinate[T] {
	return OrdinalToCoordinate(t.Rows, t.Cols, ord)
}

func (t MatrixDescriptor[T]) CoordinateToOrdinal(coord MatrixCoordinate[T]) T {
	return CoordinateToOrdinal(t.Rows, t.Cols, coord)
}
