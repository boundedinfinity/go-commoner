package slicer

func Tail[V any](items ...V) ([]V, bool) {
	if len(items) > 1 {
		return items[1:], true
	} else {
		var zero []V
		return zero, false
	}
}
