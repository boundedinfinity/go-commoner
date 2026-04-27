package mather

import "math"

// Signbit reports whether n is negative or negative zero.
func Signbit[N Number](n N) bool {
	return math.Signbit(float64(n))
}

// IsNegative reports whether n is negative or negative zero.
func IsNegative[I Number](n I) bool {
	return Signbit(n)
}
