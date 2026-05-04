package slicer

func Tail[T any](elems ...T) []T {
	if len(elems) <= 1 {
		zero := []T{}
		return zero
	}

	return elems[1:]
}

func TailOk[T any](elems ...T) ([]T, bool) {
	if len(elems) <= 1 {
		zero := []T{}
		return zero, false
	}

	return elems[1:], true
}
