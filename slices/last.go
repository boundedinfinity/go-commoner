package slices

func Last[V any](vs []V) V {
	l := len(vs)

	if l > 0 {
		return vs[l-1]
	} else {
		var zero V
		return zero
	}
}

func LastV[V any](vs ...V) V {
	return Last(vs)
}
