package slicer

func FoldLeft[I any, O any](initial O, fn func(O, I) O, items ...I) O {
	curr := initial

	for _, item := range items {
		curr = fn(curr, item)
	}

	return curr
}

func FoldLeftI[I any, O any](initial O, fn func(int, O, I) O, items ...I) O {
	curr := initial

	for i, item := range items {
		curr = fn(i, curr, item)
	}

	return curr
}

func FoldLeftErr[I any, O any](initial O, fn func(O, I) (O, error), items ...I) (O, error) {
	curr := initial
	var err error

	for _, item := range items {
		curr, err = fn(curr, item)

		if err != nil {
			return curr, err
		}
	}

	return curr, nil
}

func FoldLeftErrI[I any, O any](initial O, fn func(int, O, I) (O, error), items ...I) (O, error) {
	curr := initial
	var err error

	for i, item := range items {
		curr, err = fn(i, curr, item)

		if err != nil {
			return curr, err
		}
	}

	return curr, nil
}

func FoldRight[I any, O any](initial O, fn func(O, I) O, items ...I) O {
	curr := initial

	for i := len(items) - 1; i >= 0; i-- {
		curr = fn(curr, items[i])
	}

	return curr
}

func FoldRightI[I any, O any](initial O, fn func(int, O, I) O, items ...I) O {
	curr := initial

	for i := len(items) - 1; i >= 0; i-- {
		curr = fn(i, curr, items[i])
	}

	return curr
}

func FoldRightErr[I any, O any](initial O, fn func(O, I) (O, error), items ...I) (O, error) {
	curr := initial
	var err error

	for i := len(items) - 1; i >= 0; i-- {
		if curr, err = fn(curr, items[i]); err != nil {
			break
		}
	}

	return curr, err
}

func FoldRightErrI[I any, O any](initial O, fn func(int, O, I) (O, error), items ...I) (O, error) {
	curr := initial
	var err error

	for i := len(items) - 1; i >= 0; i-- {
		if curr, err = fn(i, curr, items[i]); err != nil {
			break
		}
	}

	return curr, err
}
