package slicer

func All[T comparable](x T, vs []T) bool {
	fn := func(v T) bool {
		return v == x
	}

	return AllFn(fn, vs)
}

func AllV[T comparable](x T, vs ...T) bool {
	return All(x, vs)
}

func AllFn[T comparable](fn func(T) bool, vs []T) bool {
	for _, v := range vs {
		if !fn(v) {
			return false
		}
	}

	return true
}

func AllFnV[T comparable](fn func(T) bool, vs ...T) bool {
	return AllFn(fn, vs)
}
