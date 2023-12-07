package mather

import (
	"math"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather/internal"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/types"
	"golang.org/x/exp/constraints"
)

func Abs[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Abs)
}

func Acos[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Acos)
}

func Acosh[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Acosh)
}

func Asin[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Asin)
}

func Asinh[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Asinh)
}

func Atan[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Atan)
}

func Atan2[T types.Numbers](y, x T) T {
	return internal.DoubleToSingle[T, T](y, x, math.Atan2)
}

func Atanh[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Atanh)
}

func Cbrt[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Cbrt)
}

func Ceil[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Ceil)
}

func Copysign[T types.Numbers](f, sign T) T {
	return internal.DoubleToSingle[T, T](f, sign, math.Copysign)
}

func Cos[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Cos)
}

func Cosh[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Cosh)
}

func Dim[T types.Numbers](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Dim)
}

func Erf[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Erf)
}

func Erfc[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Erfc)
}

func Erfcinv[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Erfcinv)
}

func Erfinv[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Erfinv)
}

func Exp[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Exp)
}

func Exp2[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Exp2)
}

func Expm1[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Expm1)
}

func FMA[T types.Numbers](x, y, z T) T {
	return internal.TripleToSingle(x, y, z, math.FMA)
}

func Floor[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Floor)
}

func Gamma[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Gamma)
}

func Hypot[T types.Numbers](p, q T) T {
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

func J0[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.J0)
}

func J1[T types.Numbers](x T) T {
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

func Log[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Log)
}

func Log10[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Log10)
}

func Log1p[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Log1p)
}

func Log2[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Log2)
}

func Logb[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Logb)
}

func Max[T types.Numbers](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Max)
}

// func MaxAll[T types.Numbers](xs ...T) T {
// 	return T(0)
// }

func Min[T types.Numbers](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Min)
}

// func MinAll[T types.Numbers](xs ...T) T {
// 	return T(0)
// }

func Mod[T types.Numbers](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Mod)
}

func Modf[I types.Numbers, O types.Numbers](x I) (O, O) {
	return internal.SingleToDouble[I, O](x, math.Modf)
}

func NaN() float64 {
	return math.NaN()
}

func Nextafter[T types.Numbers](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Nextafter)
}

func Nextafter32[T types.Numbers](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Nextafter)
}

func Pow[T types.Numbers](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Pow)
}

func Square[T types.Numbers](n T) T {
	return Pow(n, 2)
}

func Qube[T types.Numbers](n T) T {
	return Pow(n, 3)
}

func Pow10[T constraints.Integer, U types.Numbers](n T) U {
	return U(math.Pow10(int(n)))
}

func Remainder[T types.Numbers](x, y T) T {
	return internal.DoubleToSingle[T, T](x, y, math.Remainder)
}

func Round[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Round)
}

func RoundToEven[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.RoundToEven)
}

func Signbit(x float64) bool {
	return math.Signbit(x)
}

func Sin[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Sin)
}

func Sincos(x float64) (sin, cos float64) {
	return math.Sincos(x)
}

func Sinh[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Sinh)
}

func Sqrt[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Sqrt)
}

func Tan[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Tan)
}

func Tanh[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Tanh)
}

func Trunc[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Trunc)
}

func Y0[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Y0)
}

func Y1[T types.Numbers](x T) T {
	return internal.SingleToSingle[T, T](x, math.Y1)
}

func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
