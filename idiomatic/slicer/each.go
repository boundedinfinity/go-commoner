package slicer

func Each[T any](fn func(T), items ...T) {
	for _, item := range items {
		fn(item)
	}
}

func EachErr[T any](fn func(T) error, items ...T) error {
	for _, item := range items {
		if err := fn(item); err != nil {
			return err
		}
	}

	return nil
}
