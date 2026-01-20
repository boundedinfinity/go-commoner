package matrix

import (
	"github.com/boundedinfinity/go-commoner/idiomatic"
)

type MatrixCoordinate[T idiomatic.Integer] struct {
	Row T
	Col T
}

func OrdinalToCoordinate[T idiomatic.Integer](rows, cols, ord T) MatrixCoordinate[T] {
	return MatrixCoordinate[T]{
		Row: ord / cols,
		Col: ord % cols,
	}
}

func CoordinateToOrdinal[T idiomatic.Integer](rows, cols T, coord MatrixCoordinate[T]) T {
	return 0
}
