package slicer

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
