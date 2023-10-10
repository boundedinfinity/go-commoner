package slicer

func Head[V any](items []V) V {
	if len(items) > 0 {
		return items[0]
	} else {
		var zero V
		return zero
	}
}

func HeadV[V any](items ...V) V {
	return Head(items)
}
