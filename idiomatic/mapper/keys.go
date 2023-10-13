package mapper

func Keys[K comparable, V any](m map[K]V) []K {
	return KeysFiltered(m, func(k K) bool {
		return true
	})
}

func KeysFiltered[K comparable, V any](m map[K]V, fn func(K) bool) []K {
	var ks []K

	for k := range m {
		if fn(k) {
			ks = append(ks, k)
		}
	}

	return ks
}
