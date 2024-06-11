package slicer

func Tail[T any](elems ...T) ([]T, bool) {
	if len(elems) > 1 {
		return elems[1:], true
	} else {
		var zero []T
		return zero, false
	}
}
