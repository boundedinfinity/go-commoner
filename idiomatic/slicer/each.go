package slicer

func EachFunc[T any](fn func(T), elems ...T) {
	for _, elem := range elems {
		fn(elem)
	}
}

func EachFuncI[T any](fn func(int, T), elems ...T) {
	for i, elem := range elems {
		fn(i, elem)
	}
}

func EachFuncErr[T any](fn func(T) error, elems ...T) error {
	for _, elem := range elems {
		if err := fn(elem); err != nil {
			return err
		}
	}

	return nil
}

func EachFuncErrI[T any](fn func(int, T) error, elems ...T) error {
	for i, elem := range elems {
		if err := fn(i, elem); err != nil {
			return err
		}
	}

	return nil
}
