package math

import (
	"math"
)

// https://byjus.com/degree-and-radian-measure-formula/
func DegreeToRadian[T Numbers](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return n * math.Pi / 180
	})
}

// https://byjus.com/degree-and-radian-measure-formula/
func RadianToDegree[T Numbers](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return n * 180 / math.Pi
	})
}

// https://byjus.com/degree-and-radian-measure-formula/
func ArcLengthRadiusToRadian[T Numbers](arcLength, radius T) T {
	return doubleToSingle(arcLength, radius, func(a, b float64) float64 {
		return a / b
	})
}
