package mapper

func Map[K comparable, V any, R any](m map[K]V, fn func(k K, v V) R) []R {
	fn2 := func(k K, v V) (R, error) {
		return fn(k, v), nil
	}

	results, _ := MapErr(m, fn2)
	return results
}

func MapErr[K comparable, V any, R any](m map[K]V, fn func(k K, v V) (R, error)) ([]R, error) {
	var results []R
	var err error

	for k, v := range m {
		result, err := fn(k, v)

		if err != nil {
			break
		}

		results = append(results, result)
	}

	return results, err
}
