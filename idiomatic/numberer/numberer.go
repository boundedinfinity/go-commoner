package numberer

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/types"
)

func Range[T types.Number](start, end T) []T {
	return SteppedRange(start, end, 1)
}

func SteppedRange[T types.Number](start, end, step T) []T {
	var elems []T

	for i := start; i <= end; i += step {
		elems = append(elems, i)
	}

	return elems
}
