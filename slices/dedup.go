package slices

func Dedup[T comparable](items []T) []T {
	return DedupFn(items, func(item T) T {
		return item
	})
}

func DedupFn[T any, K comparable](items []T, fn func(T) K) []T {
	var o []T
	m := make(map[K]bool)

	for _, item := range items {
		key := fn(item)

		if ok := m[key]; !ok {
			m[key] = true
			o = append(o, item)
		}
	}

	return o
}
