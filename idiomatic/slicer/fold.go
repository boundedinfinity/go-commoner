package slicer

func FoldLeft[I any, O any](initial O, items []I, fn func(O, I) O) O {
	curr := initial

	for _, item := range items {
		curr = fn(curr, item)
	}

	return curr
}

func FoldIndexed[I any, O any](initial O, items []I, fn func(int, O, I) O) O {
	return FoldIndexed(initial, items, fn)
}

func FoldLeftIndexed[I any, O any](initial O, items []I, fn func(int, O, I) O) O {
	curr := initial

	for i, item := range items {
		curr = fn(i, curr, item)
	}

	return curr
}

func FoldErr[I any, O any](initial O, items []I, fn func(O, I) (O, error)) (O, error) {
	return FoldLeftErr(initial, items, fn)
}

func FoldLeftErr[I any, O any](initial O, items []I, fn func(O, I) (O, error)) (O, error) {
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

func FoldRight[I any, O any](initial O, items []I, fn func(O, I) O) O {
	curr := initial

	for i := len(items) - 1; i >= 0; i-- {
		curr = fn(curr, items[i])
	}

	return curr
}

func FoldRightIndexed[I any, O any](initial O, items []I, fn func(int, O, I) O) O {
	curr := initial

	for i := len(items) - 1; i >= 0; i-- {
		curr = fn(i, curr, items[i])
	}

	return curr
}

func FoldRightErr[I any, O any](initial O, items []I, fn func(O, I) (O, error)) (O, error) {
	curr := initial
	var err error

	for i := len(items) - 1; i >= 0; i-- {
		curr, err = fn(curr, items[i])

		if err != nil {
			return curr, err
		}
	}

	return curr, nil
}
