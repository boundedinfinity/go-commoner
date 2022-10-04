package slices

func Tail[V any](items []V) []V {
	if len(items) > 1 {
		return items[1:]
	} else {
		return []V{}
	}
}

func TailV[V any](items ...V) []V {
	return Tail(items)
}
