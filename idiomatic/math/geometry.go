package math

import "golang.org/x/exp/constraints"

type CartesianCoordinate[T constraints.Float | constraints.Integer] struct {
	X T
	Y T
}

type PolarCoordinate[T constraints.Float | constraints.Integer] struct {
	Radius T
	Angle  T
}

func CartesianToPolar[T constraints.Float | constraints.Integer](coordinate CartesianCoordinate[T]) PolarCoordinate[T] {
	return PolarCoordinate[T]{
		Radius: Sqrt(Pow(coordinate.X, 2) + Pow(coordinate.Y, 2)),
		Angle:  Atan(coordinate.Y / coordinate.X),
	}
}

func PolarToCartesian[T constraints.Float | constraints.Integer](coordinate PolarCoordinate[T]) CartesianCoordinate[T] {
	return CartesianCoordinate[T]{
		X: coordinate.Radius * Cos(coordinate.Angle),
		Y: coordinate.Radius * Sin(coordinate.Angle),
	}
}
