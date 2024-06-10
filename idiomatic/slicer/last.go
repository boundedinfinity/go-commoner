package slicer

func End[T any](items ...T) (T, bool) {
	return Last(items...)
}

func Last[T any](items ...T) (T, bool) {
	l := len(items)

	if l > 0 {
		return items[l-1], true
	} else {
		var zero T
		return zero, false
	}
}
