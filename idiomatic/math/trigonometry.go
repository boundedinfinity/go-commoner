package math

import (
	"math"

	"golang.org/x/exp/constraints"
)

// https://byjus.com/degree-and-radian-measure-formula/
func DegreeToRadian[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return n * math.Pi / 180
	})
}

// https://byjus.com/degree-and-radian-measure-formula/
func RadianToDegree[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return n * 180 / math.Pi
	})
}

// https://byjus.com/degree-and-radian-measure-formula/
func ArcLengthRadiusToRadian[T constraints.Float | constraints.Integer](arcLength, radius T) T {
	return doubleToSingle(arcLength, radius, func(a, b float64) float64 {
		return a / b
	})
}
