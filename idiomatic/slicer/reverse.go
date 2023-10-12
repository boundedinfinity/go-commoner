package slicer

func Reverse[T any](items ...T) []T {
	var reversed []T

	for i := len(items) - 1; i >= 0; i-- {
		reversed = append(reversed, items[i])
	}

	return reversed
}
