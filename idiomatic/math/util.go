package math

import "golang.org/x/exp/constraints"

type Numbers interface {
	constraints.Float | constraints.Integer
}

func singleToSingle[T Numbers](x T, fn func(float64) float64) T {
	return T(fn(float64(x)))
}

func doubleToSingle[T Numbers](a, b T, fn func(float64, float64) float64) T {
	return T(fn(float64(a), float64(b)))
}

func tripleToSingle[T Numbers](a, b, c T, fn func(float64, float64, float64) float64) T {
	return T(fn(float64(a), float64(b), float64(c)))
}

func singleToDouble[I Numbers, O Numbers](x I, fn func(float64) (float64, float64)) (O, O) {
	a, b := fn(float64(x))
	return O(a), O(b)
}
