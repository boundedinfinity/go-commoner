package math

import (
	"github.com/boundedinfinity/go-commoner/math"
	"golang.org/x/exp/constraints"
)

type CartesianCoordinate[T constraints.Float | constraints.Integer] struct {
	X T
	Y T
}

type PolarCoordinate[T constraints.Float | constraints.Integer] struct {
	Radius    T
	Angle     T
	AngleType math.AngleType
}

func CartesianToPolar[T constraints.Float | constraints.Integer](coordinate CartesianCoordinate[T]) PolarCoordinate[T] {
	return PolarCoordinate[T]{
		Radius:    Sqrt(Pow(coordinate.X, 2) + Pow(coordinate.Y, 2)),
		Angle:     Atan(coordinate.Y / coordinate.X),
		AngleType: math.AngleTypes.Radians,
	}
}

func PolarToCartesian[T constraints.Float | constraints.Integer](coordinate PolarCoordinate[T]) CartesianCoordinate[T] {
	radians := coordinate.Angle

	switch coordinate.AngleType {
	case math.AngleTypes.Radians:
		// nothing to do
	default:
		radians = DegreeToRadian(radians)
	}

	return CartesianCoordinate[T]{
		X: coordinate.Radius * Cos(radians),
		Y: coordinate.Radius * Sin(radians),
	}
}
