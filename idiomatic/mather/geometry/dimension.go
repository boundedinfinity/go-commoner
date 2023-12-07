package geometry

import "github.com/boundedinfinity/go-commoner/idiomatic/mather/types"

type Dimension2d[T types.Numbers] struct {
	Height T
	Width  T
}

func NewDimension2d[T types.Numbers](height, width T) Dimension2d[T] {
	return Dimension2d[T]{
		Height: height,
		Width:  width,
	}
}
