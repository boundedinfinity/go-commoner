package slicer

func FirstNotZero[T comparable](elems ...T) (T, bool) {
	var zero T

	return FindFn(func(elem T) bool {
		return elem != zero
	}, elems...)
}

func FirstZero[T comparable](elems ...T) (T, bool) {
	var zero T

	return FindFn(func(elem T) bool {
		return elem == zero
	}, elems...)
}
