package slicer

func Head[T any](elems ...T) (T, bool) {
	return First(elems...)
}

func First[T any](elems ...T) (T, bool) {
	if len(elems) > 0 {
		return elems[0], true
	} else {
		var zero T
		return zero, false
	}
}
