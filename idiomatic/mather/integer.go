package mather

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// https://en.wikipedia.org/wiki/Euclidean_algorithm
// https://www.youtube.com/watch?v=yHwneN6zJmU&t=641

func GreatestCommonFactor[T constraints.Integer](a T, b T) T {
	if b == 0 {
		return a
	} else {
		return GreatestCommonFactor(b, Mod(a, b))
	}
}

func AllFactors[T constraints.Integer](n T) []T {
	var factors []T
	start := T(1)

	for i := start; i <= n; i++ {
		mod := n % i

		if mod == 0 {
			factors = append(factors, i)
		}
	}

	return factors
}

func Magnitude[T constraints.Integer](n T) int {
	s := fmt.Sprintf("%v", n)
	return len(s)
}
