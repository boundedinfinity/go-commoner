package numberer

import "golang.org/x/exp/constraints"

type number interface {
	constraints.Integer | constraints.Float
}

func Range[T number](start, end T) []T {
	var items []T

	for i := start; i <= end; i++ {
		items = append(items, i)
	}

	return items
}

func RangeStep[T number](start, end, step T) []T {
	var items []T
	s := T(0)

	for i := start; i <= end; i++ {
		if i == s {
			items = append(items, i)
			s += step
		}
	}

	return items
}
