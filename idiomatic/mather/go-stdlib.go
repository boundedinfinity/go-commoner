package mather

import (
	"math"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather/internal"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/types"
)

// Abs returns the absolute value of x.
//
// Special cases are:
//
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Abs)
}

func Cbrt[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Cbrt)
}

func Ceil[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Ceil)
}

func Copysign[T types.Number](f, sign T) T {
	return internal.DoubleToSingle[T, T](f, sign, math.Copysign)
}

func Dim[T types.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Dim)
}

func Erf[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Erf)
}

func Erfc[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Erfc)
}

func Erfcinv[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Erfcinv)
}

func Erfinv[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Erfinv)
}

func Exp[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Exp)
}

func Exp2[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Exp2)
}

func Expm1[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Expm1)
}

func FMA[T types.Number](x, y, z T) T {
	return internal.TripleToSingle(x, y, z, math.FMA)
}

func Floor[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Floor)
}

func Gamma[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Gamma)
}

func Hypot[T types.Number](p, q T) T {
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

func J0[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.J0)
}

func J1[T types.Number](x T) T {
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

func Log[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Log)
}

func Log10[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Log10)
}

func Log1p[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Log1p)
}

func Log2[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Log2)
}

func Logb[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Logb)
}

func Max[T types.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Max)
}

// func MaxAll[T types.Numbers](xs ...T) T {
// 	return T(0)
// }

func Min[T types.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Min)
}

// func MinAll[T types.Numbers](xs ...T) T {
// 	return T(0)
// }

func Mod[T types.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Mod)
}

func Modf[I types.Number, O types.Number](x I) (O, O) {
	return internal.SingleToDouble[I, O](x, math.Modf)
}

func NaN() float64 {
	return math.NaN()
}

func Nextafter[T types.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Nextafter)
}

func Nextafter32[T types.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Nextafter)
}

func Pow[T types.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Pow)
}

func Square[T types.Number](n T) T {
	return Pow(n, 2)
}

func Qube[T types.Number](n T) T {
	return Pow(n, 3)
}

func Pow10[T types.Integer, U types.Number](n T) U {
	return U(math.Pow10(int(n)))
}

func Remainder[T types.Number](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Remainder)
}

func Round[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Round)
}

func RoundToEven[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.RoundToEven)
}

func Sqrt[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Sqrt)
}

func Trunc[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Trunc)
}

func Y0[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Y0)
}

func Y1[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, math.Y1)
}

func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
