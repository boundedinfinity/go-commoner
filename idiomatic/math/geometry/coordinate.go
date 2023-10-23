package geometry

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/math"
	"github.com/boundedinfinity/go-commoner/idiomatic/math/types"
)

type CartesianCoordinate[T types.Numbers] struct {
	X T
	Y T
}

type PolarCoordinate[T types.Numbers] struct {
	Radius T
	Angle  Angle[T]
}

func CartesianToPolar[T types.Numbers](coordinate CartesianCoordinate[T]) PolarCoordinate[T] {
	return PolarCoordinate[T]{
		Radius: math.Sqrt(math.Pow(coordinate.X, 2) + math.Pow(coordinate.Y, 2)),
		Angle: Angle[T]{
			Magnitude: math.Atan(coordinate.Y / coordinate.X),
			Type:      AngleTypes.Radians,
			Direction: AngleDirections.CounterClockwise,
		},
	}
}

func PolarToCartesian[T types.Numbers](coordinate PolarCoordinate[T]) CartesianCoordinate[T] {
	angle := coordinate.Angle.ToRadian()

	return CartesianCoordinate[T]{
		X: coordinate.Radius * math.Cos(angle.Magnitude),
		Y: coordinate.Radius * math.Sin(angle.Magnitude),
	}
}
