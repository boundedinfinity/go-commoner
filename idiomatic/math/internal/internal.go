package internal

import "github.com/boundedinfinity/go-commoner/idiomatic/math/types"

func ZeroToSingle[I types.Numbers](fn func() float64) I {
	return I(fn())
}

func SingleToSingle[I types.Numbers, O types.Numbers](x I, fn func(float64) float64) O {
	return O(fn(float64(x)))
}

func DoubleToSingle[I types.Numbers, O types.Numbers](a, b I, fn func(float64, float64) float64) O {
	return O(fn(float64(a), float64(b)))
}

func TripleToSingle[T types.Numbers](a, b, c T, fn func(float64, float64, float64) float64) T {
	return T(fn(float64(a), float64(b), float64(c)))
}

func QuadrupleToSingle[T types.Numbers](a, b, c, d T, fn func(float64, float64, float64, float64) float64) T {
	return T(fn(float64(a), float64(b), float64(c), float64(d)))
}

func SingleToDouble[I types.Numbers, O types.Numbers](x I, fn func(float64) (float64, float64)) (O, O) {
	a, b := fn(float64(x))
	return O(a), O(b)
}
