package mapper

func Has[K comparable, V any](m map[K]V, k K) bool {
	if m == nil {
		return false
	}

	_, ok := m[k]
	return ok
}

func HasValueFunc[K comparable, V any](m map[K]V, fn func(V) bool) bool {
	for _, v := range m {
		if fn(v) {
			return true
		}
	}

	return false
}

func HasItemFunc[K comparable, V any](m map[K]V, fn func(K, V) bool) bool {
	for k, v := range m {
		if fn(k, v) {
			return true
		}
	}

	return false
}
