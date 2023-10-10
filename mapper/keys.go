package mapper

func KeysFiltered[K comparable, V any](m map[K]V, fn func(K) bool) []K {
	var ks []K

	if m != nil {
		for k := range m {
			if fn(k) {
				ks = append(ks, k)
			}
		}
	}

	return ks
}

func Keys[K comparable, V any](m map[K]V) []K {
	return KeysFiltered(m, func(k K) bool {
		return true
	})
}