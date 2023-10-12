package slicer

func Dedup[T comparable](items ...T) []T {
	fn := func(item T) T {
		return item
	}

	return DedupFn(fn, items...)
}

func DedupFn[T any, K comparable](fn func(T) K, items ...T) []T {
	var o []T
	m := make(map[K]bool)

	for _, item := range items {
		key := fn(item)

		if ok := m[key]; !ok {
			m[key] = true
			o = append(o, item)
		}
	}

	return o
}

func DedupFnErr[T any, K comparable](fn func(T) (K, error), items ...T) ([]T, error) {
	var o []T
	m := make(map[K]bool)

	for _, item := range items {
		key, err := fn(item)

		if err != nil {
			return o, err
		}

		if ok := m[key]; !ok {
			m[key] = true
			o = append(o, item)
		}
	}

	return o, nil
}
