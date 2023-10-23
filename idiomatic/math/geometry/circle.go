package geometry

import (
	"math"

	"github.com/boundedinfinity/go-commoner/idiomatic/math/internal"
	"github.com/boundedinfinity/go-commoner/idiomatic/math/types"
)

func NewCircle[T types.Numbers](center CartesianCoordinate[T], radius T) Circle[T] {
	return Circle[T]{
		Center: center,
		Radius: radius,
	}
}

func NewCircleXY[T types.Numbers](x, y, radius T) Circle[T] {
	return NewCircle[T](CartesianCoordinate[T]{X: x, Y: y}, radius)
}

type Circle[T types.Numbers] struct {
	Center CartesianCoordinate[T]
	Radius T
}

func (t Circle[T]) PointOnCircumference(angle Angle[T]) CartesianCoordinate[T] {
	var theta T

	fn := func(theta float64) float64 { return 360 - theta }

	switch angle.Direction {
	case AngleDirections.CounterClockwise:
		theta = internal.SingleToSingle[T, T](theta, fn)
	default:
		theta = angle.Magnitude
	}

	return CartesianCoordinate[T]{
		X: internal.TripleToSingle[T](t.Radius, theta, t.Center.X, func(radius, theta, x float64) float64 {
			return radius*math.Cos(theta) + x
		}),
		Y: internal.TripleToSingle[T](t.Radius, theta, t.Center.Y, func(radius, theta, y float64) float64 {
			return radius*math.Sin(theta) + y
		}),
	}
}
