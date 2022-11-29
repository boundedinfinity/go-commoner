package mapper

type Item[K comparable, V any] struct {
	K K
	V V
}

func Items[K comparable, V any](m map[K]V) []Item[K, V] {
	var items []Item[K, V]

	for k, v := range m {
		items = append(items, Item[K, V]{K: k, V: v})
	}

	return items
}
