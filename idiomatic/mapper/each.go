package mapper

func Each[K comparable, V any](m map[K]V, fn func(K, V)) {
	for k, v := range m {
		fn(k, v)
	}
}

func EachErr[K comparable, V any](m map[K]V, fn func(K, V) error) error {
	for k, v := range m {
		if err := fn(k, v); err != nil {
			return err
		}
	}

	return nil
}
