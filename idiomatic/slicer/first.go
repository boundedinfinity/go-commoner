package slicer

func FirstNotZero[T comparable](items ...T) (T, bool) {
	var zero T

	return FindFn(func(item T) bool {
		return item != zero
	}, items...)
}

func FirstZero[T comparable](items ...T) (T, bool) {
	var zero T

	return FindFn(func(item T) bool {
		return item == zero
	}, items...)
}
