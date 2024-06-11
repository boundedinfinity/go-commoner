package numberer

import "golang.org/x/exp/constraints"

type number interface {
	constraints.Integer | constraints.Float
}

func Range[T number](start, end T) []T {
	var elems []T

	for i := start; i <= end; i++ {
		elems = append(elems, i)
	}

	return elems
}

func RangeStep[T number](start, end, step T) []T {
	var elems []T
	s := T(0)

	for i := start; i <= end; i++ {
		if i == s {
			elems = append(elems, i)
			s += step
		}
	}

	return elems
}
