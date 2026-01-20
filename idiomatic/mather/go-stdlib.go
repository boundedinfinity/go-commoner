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
	return T(math.Abs(float64(x)))
}

func Cbrt[T idiomatic.Number](x T) T {
	return T(math.Cbrt(float64(x)))
}

func Ceil[T idiomatic.Number](x T) T {
	return T(math.Ceil(float64(x)))
}

func Copysign[T idiomatic.Number](f, sign T) T {
	return T(math.Copysign(float64(f), float64(sign)))
}

func Dim[T idiomatic.Number](x, y T) T {
	return T(math.Dim(float64(x), float64(y)))
}

func Erf[T idiomatic.Number](x T) T {
	return T(math.Erf(float64(x)))
}

func Erfc[T idiomatic.Number](x T) T {
	return T(math.Erfc(float64(x)))
}

func Erfcinv[T idiomatic.Number](x T) T {
	return T(math.Erfcinv(float64(x)))
}

func Erfinv[T idiomatic.Number](x T) T {
	return T(math.Erfinv(float64(x)))
}

func Exp[T idiomatic.Number](x T) T {
	return T(math.Exp(float64(x)))
}

func Exp2[T idiomatic.Number](x T) T {
	return T(math.Exp2(float64(x)))
}

func Expm1[T idiomatic.Number](x T) T {
	return T(math.Expm1(float64(x)))
}

func FMA[T idiomatic.Number](x, y, z T) T {
	return T(math.FMA(float64(x), float64(y), float64(z)))
}

func Floor[T idiomatic.Number](x T) T {
	return T(math.Floor(float64(x)))
}

func Gamma[T idiomatic.Number](x T) T {
	return T(math.Gamma(float64(x)))
}

func Hypot[T idiomatic.Number](p, q T) T {
	return T(math.Hypot(float64(p), float64(q)))
}

func Ilogb[F idiomatic.Number, I idiomatic.Integer](x float64) I {
	return I(math.Ilogb(float64(x)))
}

func Inf[F idiomatic.Number, I idiomatic.Integer](x I) F {
	return F(math.Inf(int(x)))
}

func IsInf[F idiomatic.Number, I idiomatic.Integer](f F, sign I) bool {
	return math.IsInf(float64(f), int(sign))
}

func IsNaN[F idiomatic.Number](x F) bool {
	return math.IsNaN(float64(x))
}

func J0[T idiomatic.Number](x T) T {
	return T(math.J0(float64(x)))
}

func J1[T idiomatic.Number](x T) T {
	return T(math.J1(float64(x)))
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
	return T(math.Log(float64(x)))
}

func Log10[T idiomatic.Number](x T) T {
	return T(math.Log10(float64(x)))
}

func Log1p[T idiomatic.Number](x T) T {
	return T(math.Log1p(float64(x)))
}

func Log2[T idiomatic.Number](x T) T {
	return T(math.Log2(float64(x)))
}

func Logb[T idiomatic.Number](x T) T {
	return T(math.Logb(float64(x)))
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

func Modf[I idiomatic.Number](x I) (float64, float64) {
	return math.Modf(float64(x))
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
