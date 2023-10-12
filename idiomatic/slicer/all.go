package slicer

func All[T comparable](x T, vs ...T) bool {
	fn := func(v T) bool {
		return v == x
	}

	return AllFn(fn, vs...)
}

func AllFn[T comparable](fn func(T) bool, vs ...T) bool {
	for _, v := range vs {
		if !fn(v) {
			return false
		}
	}

	return true
}

func AllFnErr[T comparable](fn func(T) (bool, error), vs ...T) (bool, error) {
	ok := true
	var err error

	for _, v := range vs {
		ok, err = fn(v)

		if !ok || err != nil {
			break
		}
	}

	return ok, err
}
