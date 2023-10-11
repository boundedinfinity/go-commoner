package mapper

func Has[K comparable, V any](m map[K]V, k K) bool {
	if m == nil {
		return false
	}

	_, ok := m[k]
	return ok
}
