package slicer

func Each[T any](fn func(int, T), elems ...T) {
	fn2 := func(i int, elem T) error {
		fn(i, elem)
		return nil
	}

	EachErr(fn2, elems...)
}

func EachErr[T any](fn func(int, T) error, elems ...T) error {
	for i, elem := range elems {
		if err := fn(i, elem); err != nil {
			return err
		}
	}

	return nil
}
