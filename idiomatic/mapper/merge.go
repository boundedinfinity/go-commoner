package mapper

func MergeInto[K comparable, V any](m1, m2 map[K]V) {
	for k, v := range m2 {
		m1[k] = v
	}
}

func MergeCopy[K comparable, V any](m1, m2 map[K]V) {
	m3 := map[K]V{}

	for k, v := range m1 {
		m3[k] = v
	}

	for k, v := range m2 {
		m3[k] = v
	}
}
