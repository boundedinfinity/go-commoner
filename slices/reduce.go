package slices

import "github.com/boundedinfinity/commons/try"

func Reduce[V any, A any](initial A, vs []V, fn func(V, A) A) A {
	if len(vs) == 0 {
		return initial
	}

	c := initial

	for _, v := range vs {
		c = fn(v, c)
	}

	return c
}

func ReduceErr[V any, A any](vs []V, fn func(V, A) (A, error), s A) (A, error) {
	c := s

	for _, v := range vs {
		c, err := fn(v, c)

		if err != nil {
			return c, err
		}
	}

	return c, nil
}

func ReduceTry[V any, A any](vs []V, fn func(V, A) (A, error), initial A) try.Try[A] {
	agg := initial

	for _, v := range vs {
		a, err := fn(v, agg)
		agg = a

		if err != nil {
			return try.FailureR(agg, err)
		}
	}

	return try.Success(agg)
}
