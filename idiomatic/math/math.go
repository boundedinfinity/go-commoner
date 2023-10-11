package math

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Abs)
}

func Acos[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Acos)
}

func Acosh[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Acosh)
}

func Asin[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Asin)
}

func Asinh[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Asinh)
}

func Atan[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Atan)
}

func Atan2[T constraints.Float | constraints.Integer](y, x T) T {
	return doubleToSingle(y, x, math.Atan2)
}

func Atanh[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Atanh)
}

func Cbrt[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Cbrt)
}

func Ceil[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Ceil)
}

func Copysign[T constraints.Float | constraints.Integer](f, sign T) T {
	return doubleToSingle(f, sign, math.Copysign)
}

func Cos[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Cos)
}

func Cosh[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Cosh)
}

func Dim[T constraints.Float | constraints.Integer](x, y T) T {
	return doubleToSingle(x, y, math.Dim)
}

func Erf[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Erf)
}

func Erfc[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Erfc)
}

func Erfcinv[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Erfcinv)
}

func Erfinv[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Erfinv)
}

func Exp[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Exp)
}

func Exp2[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Exp2)
}

func Expm1[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Expm1)
}

func FMA[T constraints.Float | constraints.Integer](x, y, z T) T {
	return tripleToSingle(x, y, z, math.FMA)
}

func Floor[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Floor)
}

func Gamma[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Gamma)
}

func Hypot[T constraints.Float | constraints.Integer](p, q T) T {
	return doubleToSingle(p, q, math.Hypot)
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

func J0[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.J0)
}

func J1[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.J1)
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

func Log[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Log)
}

func Log10[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Log10)
}

func Log1p[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Log1p)
}

func Log2[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Log2)
}

func Logb[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Logb)
}

func Max[T constraints.Float | constraints.Integer](x, y T) T {
	return doubleToSingle(x, y, math.Max)
}

// func MaxAll[T constraints.Float | constraints.Integer](xs ...T) T {
// 	return T(0)
// }

func Min[T constraints.Float | constraints.Integer](x, y T) T {
	return doubleToSingle(x, y, math.Min)
}

// func MinAll[T constraints.Float | constraints.Integer](xs ...T) T {
// 	return T(0)
// }

func Mod[T constraints.Float | constraints.Integer](x, y T) T {
	return doubleToSingle(x, y, math.Mod)
}

func Modf[I constraints.Float | constraints.Integer, O constraints.Float | constraints.Integer](x I) (O, O) {
	return singleToDouble[I, O](x, math.Modf)
}

func NaN() float64 {
	return math.NaN()
}

func Nextafter[T constraints.Float | constraints.Integer](x, y T) T {
	return doubleToSingle(x, y, math.Nextafter)
}

func Nextafter32[T constraints.Float | constraints.Integer](x, y T) T {
	return doubleToSingle(x, y, math.Nextafter)
}

func Pow[T constraints.Float | constraints.Integer](x, y T) T {
	return doubleToSingle(x, y, math.Pow)
}

func Pow10(n int) float64 {
	return math.Pow10(n)
}

func Remainder[T constraints.Float | constraints.Integer](x, y T) T {
	return doubleToSingle(x, y, math.Remainder)
}

func Round[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Round)
}

func RoundToEven[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.RoundToEven)
}

func Signbit(x float64) bool {
	return math.Signbit(x)
}

func Sin[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Sin)
}

func Sincos(x float64) (sin, cos float64) {
	return math.Sincos(x)
}

func Sinh[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Sinh)
}

func Sqrt[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Sqrt)
}

func Tan[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Tan)
}

func Tanh[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Tanh)
}

func Trunc[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Trunc)
}

func Y0[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Y0)
}

func Y1[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, math.Y1)
}

func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
