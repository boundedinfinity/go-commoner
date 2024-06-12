package slicer

func Fold[I any, R any](initial R, fn func(int, R, I) R, elems ...I) R {
	return FoldLeft(initial, fn, elems...)
}

func FoldLeft[I any, R any](initial R, fn func(int, R, I) R, elems ...I) R {
	result := initial

	for i, elem := range elems {
		result = fn(i, result, elem)
	}

	return result
}

func FoldErr[I any, R any](initial R, fn func(int, R, I) (R, error), elems ...I) (R, error) {
	return FoldLeftErr(initial, fn, elems...)
}

func FoldLeftErr[I any, R any](initial R, fn func(int, R, I) (R, error), elems ...I) (R, error) {
	result := initial
	var err error

	for i, elem := range elems {
		result, err = fn(i, result, elem)

		if err != nil {
			return result, err
		}
	}

	return result, nil
}

func FoldRight[I any, R any](initial R, fn func(int, R, I) R, elems ...I) R {
	result := initial

	for i := len(elems) - 1; i >= 0; i-- {
		result = fn(i, result, elems[i])
	}

	return result
}

func FoldRightErr[I any, R any](initial R, fn func(int, R, I) (R, error), elems ...I) (R, error) {
	result := initial
	var err error

	for i := len(elems) - 1; i >= 0; i-- {
		if result, err = fn(i, result, elems[i]); err != nil {
			break
		}
	}

	return result, err
}
