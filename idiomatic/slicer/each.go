package slicer

func Each[T any](fn func(T), elems ...T) {
	for _, elem := range elems {
		fn(elem)
	}
}

func EachErr[T any](fn func(T) error, elems ...T) error {
	for _, elem := range elems {
		if err := fn(elem); err != nil {
			return err
		}
	}

	return nil
}
