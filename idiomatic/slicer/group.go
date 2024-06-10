package slicer

func GroupFn[T any, K comparable](fn func(T) K, items ...T) map[K]T {
	fn2 := func(item T) (K, error) {
		k := fn(item)
		return k, nil
	}

	results, _ := GroupFnErr(fn2, items...)

	return results
}

func GroupFnErr[T any, K comparable](fn func(T) (K, error), items ...T) (map[K]T, error) {
	groups := make(map[K]T)
	var err error
	var key K

	for _, item := range items {
		if key, err = fn(item); err != nil {
			break
		} else {
			if _, ok := groups[key]; !ok {
				groups[key] = item
			}
		}
	}

	return groups, err
}
