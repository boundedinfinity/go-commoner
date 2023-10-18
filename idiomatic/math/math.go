package math

import (
	"math"
)

func Abs[T Numbers](x T) T {
	return singleToSingle(x, math.Abs)
}

func Acos[T Numbers](x T) T {
	return singleToSingle(x, math.Acos)
}

func Acosh[T Numbers](x T) T {
	return singleToSingle(x, math.Acosh)
}

func Asin[T Numbers](x T) T {
	return singleToSingle(x, math.Asin)
}

func Asinh[T Numbers](x T) T {
	return singleToSingle(x, math.Asinh)
}

func Atan[T Numbers](x T) T {
	return singleToSingle(x, math.Atan)
}

func Atan2[T Numbers](y, x T) T {
	return doubleToSingle(y, x, math.Atan2)
}

func Atanh[T Numbers](x T) T {
	return singleToSingle(x, math.Atanh)
}

func Cbrt[T Numbers](x T) T {
	return singleToSingle(x, math.Cbrt)
}

func Ceil[T Numbers](x T) T {
	return singleToSingle(x, math.Ceil)
}

func Copysign[T Numbers](f, sign T) T {
	return doubleToSingle(f, sign, math.Copysign)
}

func Cos[T Numbers](x T) T {
	return singleToSingle(x, math.Cos)
}

func Cosh[T Numbers](x T) T {
	return singleToSingle(x, math.Cosh)
}

func Dim[T Numbers](x, y T) T {
	return doubleToSingle(x, y, math.Dim)
}

func Erf[T Numbers](x T) T {
	return singleToSingle(x, math.Erf)
}

func Erfc[T Numbers](x T) T {
	return singleToSingle(x, math.Erfc)
}

func Erfcinv[T Numbers](x T) T {
	return singleToSingle(x, math.Erfcinv)
}

func Erfinv[T Numbers](x T) T {
	return singleToSingle(x, math.Erfinv)
}

func Exp[T Numbers](x T) T {
	return singleToSingle(x, math.Exp)
}

func Exp2[T Numbers](x T) T {
	return singleToSingle(x, math.Exp2)
}

func Expm1[T Numbers](x T) T {
	return singleToSingle(x, math.Expm1)
}

func FMA[T Numbers](x, y, z T) T {
	return tripleToSingle(x, y, z, math.FMA)
}

func Floor[T Numbers](x T) T {
	return singleToSingle(x, math.Floor)
}

func Gamma[T Numbers](x T) T {
	return singleToSingle(x, math.Gamma)
}

func Hypot[T Numbers](p, q T) T {
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

func J0[T Numbers](x T) T {
	return singleToSingle(x, math.J0)
}

func J1[T Numbers](x T) T {
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

func Log[T Numbers](x T) T {
	return singleToSingle(x, math.Log)
}

func Log10[T Numbers](x T) T {
	return singleToSingle(x, math.Log10)
}

func Log1p[T Numbers](x T) T {
	return singleToSingle(x, math.Log1p)
}

func Log2[T Numbers](x T) T {
	return singleToSingle(x, math.Log2)
}

func Logb[T Numbers](x T) T {
	return singleToSingle(x, math.Logb)
}

func Max[T Numbers](x, y T) T {
	return doubleToSingle(x, y, math.Max)
}

// func MaxAll[T Numbers](xs ...T) T {
// 	return T(0)
// }

func Min[T Numbers](x, y T) T {
	return doubleToSingle(x, y, math.Min)
}

// func MinAll[T Numbers](xs ...T) T {
// 	return T(0)
// }

func Mod[T Numbers](x, y T) T {
	return doubleToSingle(x, y, math.Mod)
}

func Modf[I Numbers, O Numbers](x I) (O, O) {
	return singleToDouble[I, O](x, math.Modf)
}

func NaN() float64 {
	return math.NaN()
}

func Nextafter[T Numbers](x, y T) T {
	return doubleToSingle(x, y, math.Nextafter)
}

func Nextafter32[T Numbers](x, y T) T {
	return doubleToSingle(x, y, math.Nextafter)
}

func Pow[T Numbers](x, y T) T {
	return doubleToSingle(x, y, math.Pow)
}

func Pow10(n int) float64 {
	return math.Pow10(n)
}

func Remainder[T Numbers](x, y T) T {
	return doubleToSingle(x, y, math.Remainder)
}

func Round[T Numbers](x T) T {
	return singleToSingle(x, math.Round)
}

func RoundToEven[T Numbers](x T) T {
	return singleToSingle(x, math.RoundToEven)
}

func Signbit(x float64) bool {
	return math.Signbit(x)
}

func Sin[T Numbers](x T) T {
	return singleToSingle(x, math.Sin)
}

func Sincos(x float64) (sin, cos float64) {
	return math.Sincos(x)
}

func Sinh[T Numbers](x T) T {
	return singleToSingle(x, math.Sinh)
}

func Sqrt[T Numbers](x T) T {
	return singleToSingle(x, math.Sqrt)
}

func Tan[T Numbers](x T) T {
	return singleToSingle(x, math.Tan)
}

func Tanh[T Numbers](x T) T {
	return singleToSingle(x, math.Tanh)
}

func Trunc[T Numbers](x T) T {
	return singleToSingle(x, math.Trunc)
}

func Y0[T Numbers](x T) T {
	return singleToSingle(x, math.Y0)
}

func Y1[T Numbers](x T) T {
	return singleToSingle(x, math.Y1)
}

func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
