package mapper

func Copy[K comparable, V any](m map[K]V) map[K]V {
	result := map[K]V{}

	for k, v := range m {
		result[k] = v
	}

	return result
}
