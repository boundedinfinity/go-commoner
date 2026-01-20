package numberer

import "github.com/boundedinfinity/go-commoner/idiomatic"

func Range[T idiomatic.Number](start, end T) []T {
	return SteppedRange(start, end, 1)
}

func SteppedRange[T idiomatic.Number](start, end, step T) []T {
	var elems []T

	for i := start; i <= end; i += step {
		elems = append(elems, i)
	}

	return elems
}
