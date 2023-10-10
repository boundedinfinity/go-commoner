package slicer

func FoldLeft[I any, O any](initial O, vs []I, fn func(O, I) O) O {
	curr := initial

	for _, v := range vs {
		curr = fn(curr, v)
	}

	return curr
}

func FoldIndexed[I any, O any](initial O, vs []I, fn func(int, O, I) O) O {
	return FoldIndexed(initial, vs, fn)
}

func FoldLeftIndexed[I any, O any](initial O, vs []I, fn func(int, O, I) O) O {
	curr := initial

	for i, v := range vs {
		curr = fn(i, curr, v)
	}

	return curr
}

func FoldErr[I any, O any](initial O, vs []I, fn func(O, I) (O, error)) (O, error) {
	return FoldLeftErr(initial, vs, fn)
}

func FoldLeftErr[I any, O any](initial O, vs []I, fn func(O, I) (O, error)) (O, error) {
	curr := initial
	var err error

	for _, v := range vs {
		curr, err = fn(curr, v)

		if err != nil {
			return curr, err
		}
	}

	return curr, nil
}

func FoldRight[I any, O any](initial O, vs []I, fn func(O, I) O) O {
	curr := initial

	for i := len(vs) - 1; i >= 0; i-- {
		curr = fn(curr, vs[i])
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
