package math

import "golang.org/x/exp/constraints"

func singleToSingle[T constraints.Float | constraints.Integer](x T, fn func(float64) float64) T {
	return T(fn(float64(x)))
}

func doubleToSingle[T constraints.Float | constraints.Integer](a, b T, fn func(float64, float64) float64) T {
	return T(fn(float64(a), float64(b)))
}

func tripleToSingle[T constraints.Float | constraints.Integer](a, b, c T, fn func(float64, float64, float64) float64) T {
	return T(fn(float64(a), float64(b), float64(c)))
}
