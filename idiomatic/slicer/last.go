package slicer

func Last[V any](vs ...V) (V, bool) {
	l := len(vs)

	if l > 0 {
		return vs[l-1], true
	} else {
		var zero V
		return zero, false
	}
}
