package matrix

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
	"golang.org/x/exp/constraints"
)

type MatrixCoordinate[T constraints.Integer] struct {
	Row T
	Col T
}

func OrdinalToCoordinate[T constraints.Integer](rows, cols, ord T) MatrixCoordinate[T] {
	return MatrixCoordinate[T]{
		Row: mather.Floor(ord / cols),
		Col: ord % cols,
	}
}

func CoordinateToOrdinal[T constraints.Integer](rows, cols T, coord MatrixCoordinate[T]) T {
	return 0
}
