package stringer

func Chain[S ~string](s S, fns ...func(S) S) S {
	for _, fn := range fns {
		s = fn(s)
	}

	return s
}
