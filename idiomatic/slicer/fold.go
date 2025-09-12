package slicer

func Fold[T any, R any](fn func(R, T) R, initial R, elems ...T) R {
	acc := initial

	for _, elem := range elems {
		acc = fn(acc, elem)
	}

	return acc
}

func FoldI[T any, R any](fn func(int, R, T) R, initial R, elems ...T) R {
	acc := initial

	for i, elem := range elems {
		acc = fn(i, acc, elem)
	}

	return acc
}

func FoldErr[T any, R any](fn func(R, T) (R, error), initial R, elems ...T) (R, error) {
	acc := initial
	var err error

	for _, elem := range elems {
		acc, err = fn(acc, elem)

		if err != nil {
			break
		}
	}

	return acc, err
}

func FoldErrI[T any, R any](fn func(int, R, T) (R, error), initial R, elems ...T) (R, error) {
	acc := initial
	var err error

	for i, elem := range elems {
		acc, err = fn(i, acc, elem)

		if err != nil {
			break
		}
	}

	return acc, err
}

func FoldRight[T any, R any](fn func(R, T) R, initial R, elems ...T) R {
	acc := initial

	for i := len(elems) - 1; i >= 0; i-- {
		acc = fn(acc, elems[i])
	}

	return acc
}

func FoldRightI[T any, R any](fn func(int, R, T) R, initial R, elems ...T) R {
	acc := initial

	for i := len(elems) - 1; i >= 0; i-- {
		acc = fn(i, acc, elems[i])
	}

	return acc
}

func FoldRightErr[T any, R any](fn func(R, T) (R, error), initial R, elems ...T) (R, error) {
	acc := initial
	var err error

	for i := len(elems) - 1; i >= 0; i-- {
		acc, err = fn(acc, elems[i])

		if err != nil {
			break
		}
	}

	return acc, err
}

func FoldRightErrI[T any, R any](fn func(int, R, T) (R, error), initial R, elems ...T) (R, error) {
	acc := initial
	var err error

	for i := len(elems) - 1; i >= 0; i-- {
		acc, err = fn(i, acc, elems[i])

		if err != nil {
			break
		}
	}

	return acc, err
}
