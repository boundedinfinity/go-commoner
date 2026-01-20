package mather

import (
	"math"

	"github.com/boundedinfinity/go-commoner/idiomatic"
	"github.com/boundedinfinity/go-commoner/idiomatic/internal"
)

// Abs returns the absolute value of x.
//
// Special cases are:
//
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Abs)
}

func Cbrt[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Cbrt)
}

func Ceil[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Ceil)
}

func Copysign[T idiomatic.Number](f, sign T) T {
	return internal.DoubleToSingle[T, T](f, sign, math.Copysign)
}

func Dim[T idiomatic.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Dim)
}

func Erf[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Erf)
}

func Erfc[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Erfc)
}

func Erfcinv[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Erfcinv)
}

func Erfinv[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Erfinv)
}

func Exp[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Exp)
}

func Exp2[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Exp2)
}

func Expm1[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Expm1)
}

func FMA[T idiomatic.Number](x, y, z T) T {
	return internal.TripleToSingle(x, y, z, math.FMA)
}

func Floor[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Floor)
}

func Gamma[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Gamma)
}

func Hypot[T idiomatic.Number](p, q T) T {
	return internal.DoubleToSingle[T, T](p, q, math.Hypot)
}

func Ilogb(x float64) int {
	return math.Ilogb(x)
}

func Inf(x int) float64 {
	return math.Inf(x)
}

func IsInf(f float64, sign int) bool {
	return math.IsInf(f, sign)
}

func IsNaN(x float64) bool {
	return math.IsNaN(x)
}

func J0[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.J0)
}

func J1[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.J1)
}

func Jn(n int, x float64) float64 {
	return math.Jn(n, x)
}

func Ldexp(n float64, x int) float64 {
	return math.Ldexp(n, x)
}

func Lgamma(x float64) (lgamma float64, sign int) {
	return math.Lgamma(x)
}

func Log[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Log)
}

func Log10[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Log10)
}

func Log1p[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Log1p)
}

func Log2[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Log2)
}

func Logb[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Logb)
}

func Max[T idiomatic.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Max)
}

// func MaxAll[T idiomatic.Numbers](xs ...T) T {
// 	return T(0)
// }

func Min[T idiomatic.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Min)
}

// func MinAll[T idiomatic.Numbers](xs ...T) T {
// 	return T(0)
// }

func Mod[T idiomatic.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Mod)
}

func Modf[I idiomatic.Number, O idiomatic.Number](x I) (O, O) {
	return internal.SingleToDouble[I, O](x, math.Modf)
}

func NaN() float64 {
	return math.NaN()
}

func Nextafter[T idiomatic.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Nextafter)
}

func Nextafter32[T idiomatic.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Nextafter)
}

func Pow[T idiomatic.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Pow)
}

func Square[T idiomatic.Number](n T) T {
	return Pow(n, 2)
}

func Qube[T idiomatic.Number](n T) T {
	return Pow(n, 3)
}

func Pow10[T idiomatic.Integer, U idiomatic.Number](n T) U {
	return U(math.Pow10(int(n)))
}

func Remainder[T idiomatic.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Remainder)
}

func Round[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Round)
}

func RoundToEven[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.RoundToEven)
}

func Sqrt[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Sqrt)
}

func Trunc[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Trunc)
}

func Y0[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Y0)
}

func Y1[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Y1)
}

func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
