package geometry

import (
	gomath "math"

	"github.com/boundedinfinity/go-commoner/idiomatic/math"
	"github.com/boundedinfinity/go-commoner/idiomatic/math/internal"
	"github.com/boundedinfinity/go-commoner/idiomatic/math/types"
)

type CartesianCoordinate[T types.Numbers] struct {
	X T
	Y T
}

type PolarCoordinate[T types.Numbers] struct {
	Radius    T
	Angle     T
	AngleType AngleType
}

func CartesianToPolar[T types.Numbers](coordinate CartesianCoordinate[T]) PolarCoordinate[T] {
	return PolarCoordinate[T]{
		Radius:    math.Sqrt(math.Pow(coordinate.X, 2) + math.Pow(coordinate.Y, 2)),
		Angle:     math.Atan(coordinate.Y / coordinate.X),
		AngleType: AngleTypes.Radians,
	}
}

func PolarToCartesian[T types.Numbers](coordinate PolarCoordinate[T]) CartesianCoordinate[T] {
	radians := coordinate.Angle

	switch coordinate.AngleType {
	case AngleTypes.Radians:
		// nothing to do
	default:
		radians = math.DegreeToRadian(radians)
	}

	return CartesianCoordinate[T]{
		X: coordinate.Radius * math.Cos(radians),
		Y: coordinate.Radius * math.Sin(radians),
	}
}

func DegreeToRadian[T types.Numbers](angle T) T {
	return internal.SingleToSingle[T, T](angle, func(n float64) float64 {
		return n * gomath.Pi / 180
	})
}

func RadianToDegree[T types.Numbers](angle T) T {
	return internal.SingleToSingle[T, T](angle, func(n float64) float64 {
		return n * 180 / gomath.Pi
	})
}
