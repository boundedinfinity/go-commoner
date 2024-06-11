package slicer

func Reverse[T any](elems ...T) []T {
	var reversed []T

	for i := len(elems) - 1; i >= 0; i-- {
		reversed = append(reversed, elems[i])
	}

	return reversed
}
