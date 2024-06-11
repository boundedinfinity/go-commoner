package slicer

func Fold[I any, O any](initial O, fn func(int, O, I) O, elems ...I) O {
	return FoldLeft(initial, fn, elems...)
}

func FoldLeft[I any, O any](initial O, fn func(int, O, I) O, elems ...I) O {
	curr := initial

	for i, elem := range elems {
		curr = fn(i, curr, elem)
	}

	return curr
}

func FoldErr[I any, O any](initial O, fn func(int, O, I) (O, error), elems ...I) (O, error) {
	return FoldLeftErr(initial, fn, elems...)
}

func FoldLeftErr[I any, O any](initial O, fn func(int, O, I) (O, error), elems ...I) (O, error) {
	curr := initial
	var err error

	for i, elem := range elems {
		curr, err = fn(i, curr, elem)

		if err != nil {
			return curr, err
		}
	}

	return curr, nil
}

func FoldRight[I any, O any](initial O, fn func(int, O, I) O, elems ...I) O {
	curr := initial

	for i := len(elems) - 1; i >= 0; i-- {
		curr = fn(i, curr, elems[i])
	}

	return curr
}

func FoldRightErr[I any, O any](initial O, fn func(int, O, I) (O, error), elems ...I) (O, error) {
	curr := initial
	var err error

	for i := len(elems) - 1; i >= 0; i-- {
		if curr, err = fn(i, curr, elems[i]); err != nil {
			break
		}
	}

	return curr, err
}
