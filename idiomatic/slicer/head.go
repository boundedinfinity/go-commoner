package slicer

func Head[V any](items ...V) (V, bool) {
	if len(items) > 0 {
		return items[0], true
	} else {
		var zero V
		return zero, false
	}
}
