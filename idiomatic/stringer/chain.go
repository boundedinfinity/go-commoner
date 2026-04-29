package stringer

func Chain[S ~string](s S, fns ...func(S) S) S {
	for _, fn := range fns {
		s = fn(s)
	}

	return s
}

func ChainErr[S ~string](s S, fns ...func(S) (S, error)) (S, error) {
	var err error

	for _, fn := range fns {
		s, err = fn(s)
		if err != nil {
			return s, err
		}
	}

	return s, nil
}
