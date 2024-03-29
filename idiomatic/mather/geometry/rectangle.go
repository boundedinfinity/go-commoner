package geometry

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/internal"
)

func NewRectangle[T geometryNumber](topLeft CartesianCoordinate[T], dimensions Dimension2d[T]) Rectangle[T] {
	return Rectangle[T]{
		TopLeft:    topLeft,
		Dimensions: dimensions,
	}
}

func NewRectangleXYHW[T geometryNumber](x, y, height, width T) Rectangle[T] {
	return NewRectangle[T](
		CartesianCoordinate[T]{X: x, Y: y},
		Dimension2d[T]{Height: height, Width: width},
	)
}

type Rectangle[T geometryNumber] struct {
	TopLeft    CartesianCoordinate[T]
	Dimensions Dimension2d[T]
}

func (t Rectangle[T]) Area() T {
	return t.Dimensions.Height * t.Dimensions.Width
}

func (t Rectangle[T]) Perimeter() T {
	fn := func(height, width float64) float64 {
		return 2 * (height + width)
	}

	return internal.DoubleToSingle[T, T](t.Dimensions.Height, t.Dimensions.Width, fn)
}

func (t Rectangle[T]) BottomLeft() CartesianCoordinate[T] {
	return CartesianCoordinate[T]{
		X: t.TopLeft.X,
		Y: t.TopLeft.Y + t.Dimensions.Height,
	}
}

func (t Rectangle[T]) TopRight() CartesianCoordinate[T] {
	return CartesianCoordinate[T]{
		X: t.TopLeft.X + t.Dimensions.Width,
		Y: t.TopLeft.Y,
	}
}

func (t Rectangle[T]) BottomRight() CartesianCoordinate[T] {
	return CartesianCoordinate[T]{
		X: t.TopLeft.X + t.Dimensions.Width,
		Y: t.TopLeft.Y + t.Dimensions.Height,
	}
}

func (t Rectangle[T]) TopMidpoint() CartesianCoordinate[T] {
	return NewLineSegmentCoords(t.TopLeft, t.BottomRight()).Midpoint()
}

func (t Rectangle[T]) RightMidpoint() CartesianCoordinate[T] {
	return NewLineSegmentCoords(t.TopRight(), t.BottomRight()).Midpoint()
}

func (t Rectangle[T]) BottomMidpoint() CartesianCoordinate[T] {
	return NewLineSegmentCoords(t.BottomLeft(), t.BottomRight()).Midpoint()
}

func (t Rectangle[T]) LeftMidpoint() CartesianCoordinate[T] {
	return NewLineSegmentCoords(t.TopLeft, t.BottomLeft()).Midpoint()
}

func (t Rectangle[T]) Center() CartesianCoordinate[T] {
	return NewLineSegmentCoords(t.TopMidpoint(), t.BottomMidpoint()).Midpoint()
}

func (t Rectangle[T]) PointOnPeremiter(angle Angle[T]) CartesianCoordinate[T] {
	center := t.Center()
	centerToTopLeft := NewLineSegmentCoords(center, t.TopLeft)
	circle := NewCircle(center, centerToTopLeft.Length())
	circlePoint := circle.PointOnCircumference(angle)
	circleLine := NewLineSegmentCoords(center, circlePoint)
	topLine := NewLineSegmentCoords(t.TopLeft, t.TopRight())

	return topLine.IntersectAt(circleLine)
}
