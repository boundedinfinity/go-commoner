package mather

func Range[T Number](start, end T) []T {
	return SteppedRange(start, end, 1)
}

func SteppedRange[T Number](start, end, step T) []T {
	var elems []T

	for i := start; i <= end; i += step {
		elems = append(elems, i)
	}

	return elems
}
