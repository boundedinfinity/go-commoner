package mapper

func Clear[K comparable, V any](m *map[K]V) {
	*m = map[K]V{}
}
