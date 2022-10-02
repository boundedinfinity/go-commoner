package slices

import (
	"github.com/boundedinfinity/go-commoner/try"
)

func Fold[I any, O any](initial O, vs []I, fn func(O, I) O) O {
	return FoldLeft(initial, vs, fn)
}

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

func FoldTry[I any, O any](initial O, vs []I, fn func(O, I) (O, error)) try.Try[O] {
	return FoldLeftTry(initial, vs, fn)
}

func FoldLeftTry[I any, O any](initial O, vs []I, fn func(O, I) (O, error)) try.Try[O] {
	curr, err := FoldLeftErr(initial, vs, fn)

	if err != nil {
		try.Complete(curr, err)
	}

	return try.Success(curr)
}

func FoldRight[I any, O any](initial O, vs []I, fn func(O, I) O) O {
	curr := initial

	for i := len(vs) - 1; i >= 0; i-- {
		curr = fn(curr, vs[i])
	}

	return curr
}

func FoldRightIndexed[I any, O any](initial O, vs []I, fn func(int, O, I) O) O {
	curr := initial

	for i := len(vs) - 1; i >= 0; i-- {
		curr = fn(i, curr, vs[i])
	}

	return curr
}

func FoldRightErr[I any, O any](initial O, vs []I, fn func(O, I) (O, error)) (O, error) {
	curr := initial
	var err error

	for i := len(vs) - 1; i >= 0; i-- {
		curr, err = fn(curr, vs[i])

		if err != nil {
			return curr, err
		}
	}

	return curr, nil
}

func FoldRightTry[I any, O any](initial O, vs []I, fn func(O, I) (O, error)) try.Try[O] {
	curr, err := FoldRightErr(initial, vs, fn)
	return try.Complete(curr, err)
}
