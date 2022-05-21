package slices

func Each[T any](ts []T, fn func(T)) {
	for _, t := range ts {
		fn(t)
	}
}

func EachErr[T any](ts []T, fn func(T) error) error {
	for _, t := range ts {
		if err := fn(t); err != nil {
			return err
		}
	}

	return nil
}
