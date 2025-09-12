package mapper

func Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	if fn == nil {
		return m
	}

	results := map[K]V{}

	for k, v := range m {
		if fn(k, v) {
			results[k] = v
		}
	}

	return results
}

func FilterErr[K comparable, V any](m map[K]V, fn func(K, V) (bool, error)) (map[K]V, error) {
	if fn == nil {
		return m, nil
	}

	results := map[K]V{}

	for k, v := range m {
		ok, err := fn(k, v)

		if err != nil {
			return results, &MapperErrDetails[K, V]{
				Key:    k,
				Val:    v,
				Reason: err,
			}
		}

		if ok {
			results[k] = v
		}
	}

	return results, nil
}
