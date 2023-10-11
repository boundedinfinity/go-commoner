package math

import "golang.org/x/exp/constraints"

// https://en.wikipedia.org/wiki/Euclidean_algorithm
// https://www.youtube.com/watch?v=yHwneN6zJmU&t=641

func GreatestCommonDivisor[T constraints.Float | constraints.Integer](a T, b T) T {
	if b == 0 {
		return a
	} else {
		return GreatestCommonDivisor[T](b, Mod(a, b))
	}
}
