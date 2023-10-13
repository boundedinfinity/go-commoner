package slicer

func Dedup[T comparable](items ...T) []T {
	fn := func(item T) T {
		return item
	}

	return DedupFn(fn, items...)
}

func DedupFn[T any, K comparable](fn func(T) K, items ...T) []T {
	fn2 := func(item T) (K, error) {
		k := fn(item)
		return k, nil
	}

	results, _ := DedupFnErr(fn2, items...)

	return results
}

func DedupFnErr[T any, K comparable](fn func(T) (K, error), items ...T) ([]T, error) {
	m := make(map[K]T)
	var err error
	var key K

	for _, item := range items {
		if key, err = fn(item); err != nil {
			break
		} else {
			if _, ok := m[key]; !ok {
				m[key] = item
			}
		}
	}

	var results []T

	if err != nil {
		return results, err
	}

	for _, item := range m {
		results = append(results, item)
	}

	return results, nil
}
