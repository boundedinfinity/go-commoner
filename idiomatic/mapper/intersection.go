package mapper

func Intersection[K comparable, V any](m1, m2 map[K]V) map[K]V {
	result := map[K]V{}

	for k := range m1 {
		if v2, ok := m2[k]; ok {
			result[k] = v2
		}
	}

	return result
}
