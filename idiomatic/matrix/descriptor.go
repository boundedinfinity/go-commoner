package matrix

import "golang.org/x/exp/constraints"

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
