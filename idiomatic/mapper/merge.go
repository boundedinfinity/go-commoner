package mapper

func Merge[K comparable, V any](dst map[K]V, src map[K]V) {
	if dst == nil {
		return
	}

	for k, v := range src {
		dst[k] = v
	}
}
