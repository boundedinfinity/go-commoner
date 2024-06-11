package slicer

func End[T any](elems ...T) (T, bool) {
	return Last(elems...)
}

func Last[T any](elems ...T) (T, bool) {
	l := len(elems)

	if l > 0 {
		return elems[l-1], true
	} else {
		var zero T
		return zero, false
	}
}
