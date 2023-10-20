package matrix

import (
	"golang.org/x/exp/constraints"
)

type Matrix[T constraints.Integer, V any] struct {
	MatrixDescriptor[T]
	matrix [][]V
}

func FromDescriptor[T constraints.Integer, V any](descriptor MatrixDescriptor[T], initial V) *Matrix[T, V] {
	matrix := Matrix[T, V]{
		matrix: make([][]V, descriptor.Rows),
	}

	matrix.Fill(initial)

	return &matrix
}

func FromSlice[T constraints.Integer, V any](slice [][]V) *Matrix[T, V] {
	var zero V
	matrix := FromDescriptor[T, V](SliceToDescriptor[T, V](slice), zero)

	for row := 0; row < len(slice); row++ {
		for col := 0; col < len(slice[row]); col++ {
			matrix.matrix[row][col] = slice[row][col]
		}
	}

	return matrix
}

func (t *Matrix[T, V]) Fill(value V) {
	rows := int(t.Rows)
	cols := int(t.Cols)

	for row := 0; row < rows; row++ {
		t.matrix[0] = make([]V, t.Cols)

		for col := 0; col < cols; col++ {
			t.matrix[row][col] = value
		}
	}
}

func (t Matrix[T, V]) GetValueP(row, col T) (V, bool) {
	return t.matrix[row][col], true
}

func (t Matrix[T, V]) GetValueC(coord MatrixCoordinate[T]) (V, bool) {
	return t.GetValueP(coord.Row, coord.Col)
}
