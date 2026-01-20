package internal

import (
	"github.com/boundedinfinity/go-commoner/idiomatic"
)

// ZeroToSingle converts the returned value of fn() float64 to the idiomatic.Number
func ZeroToSingle[O idiomatic.Number](fn func() float64) O {
	return O(fn())
}

// SingleToSingle converts x of type I to a float64 to pass as a parameter to fn(float64) float64,
// then converts the returned value to type O
func SingleToSingle[I idiomatic.Number, O idiomatic.Number](x I, fn func(float64) float64) O {
	return O(fn(float64(x)))
}

// DoubleToSingle converts a and b of type I to a float64 to pass as a parameters to fn(float64, float64) float64,
// then converts the returned value to type O
func DoubleToSingle[I idiomatic.Number, O idiomatic.Number](a, b I, fn func(float64, float64) float64) O {
	return O(fn(float64(a), float64(b)))
}

// DoubleToSingle converts a, b and c of type I to a float64 to pass as a parameters to
// fn(float64, float64, float64) float64, then converts the returned value to type O
func TripleToSingle[T idiomatic.Number](a, b, c T, fn func(float64, float64, float64) float64) T {
	return T(fn(float64(a), float64(b), float64(c)))
}

// QuadrupleToSingle converts a, b, c and d of type I to a float64 to pass as a parameters to
// fn(float64, float64, float64, float64) float64, then converts the returned value to type O
func QuadrupleToSingle[T idiomatic.Number](a, b, c, d T, fn func(float64, float64, float64, float64) float64) T {
	return T(fn(float64(a), float64(b), float64(c), float64(d)))
}

// SingleToDouble converts x of type I to a float64 to pass as a parameter to fn(float64) (float64, float64),
// then converts the 2 return values to type O
func SingleToDouble[I idiomatic.Number, O idiomatic.Number](x I, fn func(float64) (float64, float64)) (O, O) {
	a, b := fn(float64(x))
	return O(a), O(b)
}
