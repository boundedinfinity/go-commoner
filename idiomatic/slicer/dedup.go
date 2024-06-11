package slicer

func Dedup[T comparable](elems ...T) []T {
	fn := func(elem T) T {
		return elem
	}

	return DedupFn(fn, elems...)
}

func DedupFn[T any, K comparable](fn func(T) K, elems ...T) []T {
	fn2 := func(elem T) (K, error) {
		k := fn(elem)
		return k, nil
	}

	results, _ := DedupFnErr(fn2, elems...)

	return results
}

func DedupFnErr[T any, K comparable](fn func(T) (K, error), elems ...T) ([]T, error) {
	m := make(map[K]T)
	var err error
	var key K
	var results []T

	for _, elem := range elems {
		if key, err = fn(elem); err != nil {
			break
		} else {
			if _, ok := m[key]; !ok {
				m[key] = elem
				results = append(results, elem)
			}
		}
	}

	if err != nil {
		return results, err
	}

	return results, nil
}
