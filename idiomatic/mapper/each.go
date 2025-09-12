package mapper

func Each[K comparable, V any](m map[K]V, fn func(K, V)) {
	if fn == nil {
		return
	}

	for k, v := range m {
		fn(k, v)
	}
}

func EachErr[K comparable, V any](m map[K]V, fn func(K, V) error) error {
	if fn == nil {
		return nil
	}

	for k, v := range m {
		if err := fn(k, v); err != nil {
			return &MapperErrDetails[K, V]{
				Key:    k,
				Val:    v,
				Reason: err,
			}
		}
	}

	return nil
}
