package slices

func Head[V any](vs ...V) V {
	if len(vs) > 0 {
		return vs[0]
	} else {
		var zero V
		return zero
	}
}

func Tail[V any](vs ...V) V {
	l := len(vs)

	if l > 0 {
		return vs[l-1]
	} else {
		var zero V
		return zero
	}
}
