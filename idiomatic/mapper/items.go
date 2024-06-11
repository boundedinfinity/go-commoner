package mapper

type Item[K comparable, V any] struct {
	K K
	V V
}

func Items[K comparable, V any](m map[K]V) []Item[K, V] {
	fn := func(k K, v V) bool { return true }
	return ItemsFiltered(m, fn)
}

func ItemsFiltered[K comparable, V any](m map[K]V, fn func(K, V) bool) []Item[K, V] {
	var elems []Item[K, V]

	for k, v := range m {
		if fn(k, v) {
			elems = append(elems, Item[K, V]{K: k, V: v})
		}
	}

	return elems
}
