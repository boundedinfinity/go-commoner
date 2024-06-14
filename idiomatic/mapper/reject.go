package mapper

func Reject[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	if fn == nil {
		return m
	}

	fn2 := func(k K, v V) (bool, error) { return fn(k, v), nil }
	results, _ := RejectErr(m, fn2)
	return results
}

func RejectErr[K comparable, V any](m map[K]V, fn func(K, V) (bool, error)) (map[K]V, error) {
	if fn == nil {
		return m, nil
	}

	results := map[K]V{}

	for k, v := range m {
		ok, err := fn(k, v)

		if err != nil {
			return results, err
		}

		if !ok {
			results[k] = v
		}
	}

	return results, nil
}
