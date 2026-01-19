package matrix

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/types"
)

type MatrixCoordinate[T types.Integer] struct {
	Row T
	Col T
}

func OrdinalToCoordinate[T types.Integer](rows, cols, ord T) MatrixCoordinate[T] {
	return MatrixCoordinate[T]{
		Row: mather.Floor(ord / cols),
		Col: ord % cols,
	}
}

func CoordinateToOrdinal[T types.Integer](rows, cols T, coord MatrixCoordinate[T]) T {
	return 0
}
