package slicer

import "github.com/boundedinfinity/go-commoner/functional/trier"

func ReduceTry[V any, A any](vs []V, fn func(V, A) (A, error), initial A) trier.Try[A] {
	agg := initial

	for _, v := range vs {
		a, err := fn(v, agg)
		agg = a

		if err != nil {
			return trier.Complete(agg, err)
		}
	}

	return trier.Success(agg)
}
