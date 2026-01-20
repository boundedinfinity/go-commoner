package trigonometry

import (
	"math"

	"github.com/boundedinfinity/go-commoner/idiomatic"
	"github.com/boundedinfinity/go-commoner/idiomatic/internal"
)

// https://byjus.com/degree-and-radian-measure-formula/
func DegreeToRadian[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return n * math.Pi / 180
	})
}

// https://byjus.com/degree-and-radian-measure-formula/
func RadianToDegree[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return n * 180 / math.Pi
	})
}

// https://byjus.com/degree-and-radian-measure-formula/
func ArcLengthRadiusToRadian[T idiomatic.Number](arcLength, radius T) T {
	return internal.DoubleToSingle[T, T](arcLength, radius, func(a, b float64) float64 {
		return a / b
	})
}

// ArcCosine returns the arccosine, in radians, of x.
//
// Special case is:
//
//	Acos(x) = NaN if x < -1 or x > 1
func ArcCosine[N idiomatic.Number](x N) N {
	return internal.SingleToSingle[N, N](x, math.Acos)
}

// InverseHyperbolicCosine returns the inverse hyperbolic cosine of x.
//
// Special cases are:
//
//	InverseHyperbolicCosine(+Inf) = +Inf
//	InverseHyperbolicCosine(x) = NaN if x < 1
//	InverseHyperbolicCosine(NaN) = NaN
func InverseHyperbolicCosine[N idiomatic.Number](x N) N {
	return internal.SingleToSingle[N, N](x, math.Acosh)
}

// ArcSine returns the arcsine, in radians, of x.
//
// Special cases are:
//
//	ArcSine(±0) = ±0
//	ArcSine(x) = NaN if x < -1 or x > 1
func ArcSine[N idiomatic.Number](x N) N {
	return internal.SingleToSingle[N, N](x, math.Asin)
}

// InverseHyperbolicSine returns the inverse hyperbolic sine of x.
//
// Special cases are:
//
//	InverseHyperbolicSine(±0) = ±0
//	InverseHyperbolicSine(±Inf) = ±Inf
//	InverseHyperbolicSine(NaN) = NaN
func InverseHyperbolicSine[N idiomatic.Number](x N) N {
	return internal.SingleToSingle[N, N](x, math.Asinh)
}

// ArcTangent returns the arctangent, in radians, of x.
//
// Special cases are:
//
//	ArcTangent(±0) = ±0
//	ArcTangent(±Inf) = ±Pi/2
func ArcTangent[N idiomatic.Number](x N) N {
	return internal.SingleToSingle[N, N](x, math.Atan)
}

// ArcTangent2 returns the arc tangent of y/x, using
// the signs of the two to determine the quadrant
// of the return value.
//
// Special cases are (in order):
//
//	ArcTangent2(y, NaN) = NaN
//	ArcTangent2(NaN, x) = NaN
//	ArcTangent2(+0, x>=0) = +0
//	ArcTangent2(-0, x>=0) = -0
//	ArcTangent2(+0, x<=-0) = +Pi
//	ArcTangent2(-0, x<=-0) = -Pi
//	ArcTangent2(y>0, 0) = +Pi/2
//	ArcTangent2(y<0, 0) = -Pi/2
//	ArcTangent2(+Inf, +Inf) = +Pi/4
//	ArcTangent2(-Inf, +Inf) = -Pi/4
//	ArcTangent2(+Inf, -Inf) = 3Pi/4
//	ArcTangent2(-Inf, -Inf) = -3Pi/4
//	ArcTangent2(y, +Inf) = 0
//	ArcTangent2(y>0, -Inf) = +Pi
//	ArcTangent2(y<0, -Inf) = -Pi
//	ArcTangent2(+Inf, x) = +Pi/2
//	ArcTangent2(-Inf, x) = -Pi/2
func ArcTangent2[N idiomatic.Number](y, x N) N {
	return internal.DoubleToSingle[N, N](y, x, math.Atan2)
}

// InverseHyperbolicTangent returns the inverse hyperbolic tangent of x.
//
// Special cases are:
//
//	InverseHyperbolicTangent(1) = +Inf
//	InverseHyperbolicTangent(±0) = ±0
//	InverseHyperbolicTangent(-1) = -Inf
//	InverseHyperbolicTangent(x) = NaN if x < -1 or x > 1
//	InverseHyperbolicTangent(NaN) = NaN
func InverseHyperbolicTangent[N idiomatic.Number](x N) N {
	return internal.SingleToSingle[N, N](x, math.Atanh)
}

// Cosine returns the cosine of the radian argument x.
//
// Special cases are:
//
//	Cosine(±Inf) = NaN
//	Cosine(NaN) = NaN
func Cosine[N idiomatic.Number](x N) N {
	return internal.SingleToSingle[N, N](x, math.Cos)
}

// HyperbolicCosine returns the hyperbolic cosine of x.
//
// Special cases are:
//
//	HyperbolicCosine(±0) = 1
//	HyperbolicCosine(±Inf) = +Inf
//	HyperbolicCosine(NaN) = NaN
func HyperbolicCosine[N idiomatic.Number](x N) N {
	return internal.SingleToSingle[N, N](x, math.Cosh)
}

// Sine returns the sine of the radian argument x.
//
// Special cases are:
//
//	Sine(±0) = ±0
//	Sine(±Inf) = NaN
//	Sine(NaN) = NaN
func Sine[N idiomatic.Number](x N) N {
	return internal.SingleToSingle[N, N](x, math.Sin)
}

// Sincos returns Sin(x), Cos(x).
//
// Special cases are:
//
//	Sincos(±0) = ±0, 1
//	Sincos(±Inf) = NaN, NaN
//	Sincos(NaN) = NaN, NaN
func Sincos[N idiomatic.Number](x N) (sin, cos N) {
	return internal.SingleToDouble[N, N](x, math.Sincos)
}

// HyperbolicSine returns the hyperbolic sine of x.
//
// Special cases are:
//
//	HyperbolicSine(±0) = ±0
//	HyperbolicSine(±Inf) = ±Inf
//	HyperbolicSine(NaN) = NaN
func HyperbolicSine[N idiomatic.Number](x N) N {
	return internal.SingleToSingle[N, N](x, math.Sinh)
}

// Tangent returns the tangent of the radian argument x.
//
// Special cases are:
//
//	Tangent(±0) = ±0
//	Tangent(±Inf) = NaN
//	Tangent(NaN) = NaN
func Tangent[N idiomatic.Number](x N) N {
	return internal.SingleToSingle[N, N](x, math.Tan)
}

// HyperbolicTan returns the hyperbolic tangent of x.
//
// Special cases are:
//
//	HyperbolicTan(±0) = ±0
//	HyperbolicTan(±Inf) = ±1
//	HyperbolicTan(NaN) = NaN
func HyperbolicTan[N idiomatic.Number](x N) N {
	return internal.SingleToSingle[N, N](x, math.Tanh)
}
