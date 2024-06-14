package mapper

func MergeInto[K comparable, V any](dst map[K]V, src map[K]V) {
	if dst == nil {
		return
	}

	for k, v := range src {
		dst[k] = v
	}
}

func MergeCopy[K comparable, V any](m1, m2 map[K]V) map[K]V {
	result := map[K]V{}

	for k, v := range m1 {
		result[k] = v
	}

	for k, v := range m2 {
		result[k] = v
	}

	return result
}
