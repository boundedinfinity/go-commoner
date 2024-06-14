package mapper

func Each[K comparable, V any](m map[K]V, fn func(K, V)) {
	if fn == nil {
		return
	}

	fn2 := func(k K, v V) error { fn(k, v); return nil }
	EachErr(m, fn2)
}

func EachErr[K comparable, V any](m map[K]V, fn func(K, V) error) error {
	if fn == nil {
		return nil
	}

	for k, v := range m {
		if err := fn(k, v); err != nil {
			return err
		}
	}

	return nil
}
