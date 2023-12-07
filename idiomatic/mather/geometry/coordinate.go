package geometry

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
)

type CartesianCoordinate[T geometryNumber] struct {
	X T
	Y T
}

type PolarCoordinate[T geometryNumber] struct {
	Radius T
	Angle  Angle[T]
}

func CartesianToPolar[T geometryNumber](coordinate CartesianCoordinate[T]) PolarCoordinate[T] {
	return PolarCoordinate[T]{
		Radius: mather.Sqrt(mather.Pow(coordinate.X, 2) + mather.Pow(coordinate.Y, 2)),
		Angle: Angle[T]{
			Magnitude: mather.Atan(coordinate.Y / coordinate.X),
			Type:      AngleTypes.Radians,
			Direction: AngleDirections.CounterClockwise,
		},
	}
}

func PolarToCartesian[T geometryNumber](coordinate PolarCoordinate[T]) CartesianCoordinate[T] {
	angle := coordinate.Angle.ToRadian()

	return CartesianCoordinate[T]{
		X: coordinate.Radius * mather.Cos(angle.Magnitude),
		Y: coordinate.Radius * mather.Sin(angle.Magnitude),
	}
}
