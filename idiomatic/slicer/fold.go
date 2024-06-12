package slicer

func Fold[T any, R any](fn func(int, R, T) R, initial R, elems ...T) R {
	return FoldLeft(fn, initial, elems...)
}

func FoldErr[T any, R any](fn func(int, R, T) (R, error), initial R, elems ...T) (R, error) {
	return FoldLeftErr(fn, initial, elems...)
}

func FoldLeft[T any, R any](fn func(int, R, T) R, initial R, elems ...T) R {
	fn2 := func(i int, accumulator R, elem T) (R, error) {
		return fn(i, accumulator, elem), nil
	}

	accumalater, _ := FoldLeftErr(fn2, initial, elems...)
	return accumalater
}

func FoldLeftErr[T any, R any](fn func(int, R, T) (R, error), initial R, elems ...T) (R, error) {
	accumulator := initial
	var err error

	for i, elem := range elems {
		accumulator, err = fn(i, accumulator, elem)

		if err != nil {
			return accumulator, err
		}
	}

	return accumulator, nil
}

func FoldRight[T any, R any](fn func(int, R, T) R, initial R, elems ...T) R {
	fn2 := func(i int, accumulator R, elem T) (R, error) {
		return fn(i, accumulator, elem), nil
	}

	accumulator, _ := FoldRightErr(fn2, initial, elems...)
	return accumulator
}

func FoldRightErr[T any, R any](fn func(int, R, T) (R, error), initial R, elems ...T) (R, error) {
	accumulator := initial
	var err error

	for i := len(elems) - 1; i >= 0; i-- {
		accumulator, err = fn(i, accumulator, elems[i])

		if err != nil {
			break
		}
	}

	return accumulator, err
}
