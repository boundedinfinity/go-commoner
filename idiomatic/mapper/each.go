package mapper

func Each[K comparable, V any](m map[K]V, fn func(K, V)) {
	fn2 := func(k K, v V) error { return nil }
	EachErr(m, fn2)
}

func EachErr[K comparable, V any](m map[K]V, fn func(K, V) error) error {
	for k, v := range m {
		if err := fn(k, v); err != nil {
			return err
		}
	}

	return nil
}
