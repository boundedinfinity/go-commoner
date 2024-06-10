package slicer

func Head[T any](items ...T) (T, bool) {
	return First(items...)
}

func First[T any](items ...T) (T, bool) {
	if len(items) > 0 {
		return items[0], true
	} else {
		var zero T
		return zero, false
	}
}
