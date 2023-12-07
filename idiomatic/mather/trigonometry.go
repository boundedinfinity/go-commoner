package mather

import (
	"math"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather/internal"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/types"
)

// https://byjus.com/degree-and-radian-measure-formula/
func DegreeToRadian[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return n * math.Pi / 180
	})
}

// https://byjus.com/degree-and-radian-measure-formula/
func RadianToDegree[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return n * 180 / math.Pi
	})
}

// https://byjus.com/degree-and-radian-measure-formula/
func ArcLengthRadiusToRadian[T types.Numbers](arcLength, radius T) T {
	return internal.DoubleToSingle[T, T](arcLength, radius, func(a, b float64) float64 {
		return a / b
	})
}
