package slicer

func Equals[T comparable](as, bs []T) bool {
	fn := func(x, y T) bool {
		return x == y
	}

	return EqualsFn(as, bs, fn)
}

func EqualsFn[T comparable](as, bs []T, fn func(T, T) bool) bool {
	l := len(as)

	if l != len(bs) {
		return false
	}

	for i := 0; i < l; i++ {
		if !fn(as[i], bs[i]) {
			return false
		}
	}

	return true
}

func EqualsFnErr[T comparable](as, bs []T, fn func(T, T) (bool, error)) (bool, error) {
	l := len(as)

	if l != len(bs) {
		return false, nil
	}

	var ok bool
	var err error

	for i := 0; i < l; i++ {
		if ok, err = fn(as[i], bs[i]); ok || err != nil {
			break
		}
	}

	return ok, err
}
