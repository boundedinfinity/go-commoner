package mapper

func Has[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]

	if ok {
		return true
	}

	return false
}
