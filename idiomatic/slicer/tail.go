package slicer

func Tail[T any](items ...T) ([]T, bool) {
	if len(items) > 1 {
		return items[1:], true
	} else {
		var zero []T
		return zero, false
	}
}
