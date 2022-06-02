package slices

import "github.com/boundedinfinity/optioner"

func Head[V any](vs []V) V {
	if len(vs) > 0 {
		return vs[0]
	} else {
		var zero V
		return zero
	}
}

func HeadV[V any](vs ...V) V {
	return Head(vs)
}

func HeadO[V any](vs []V) optioner.Optioner[V] {
	if len(vs) > 0 {
		return optioner.Some(vs[0])
	}

	return optioner.None[V]()
}

func HeadVO[V any](vs ...V) optioner.Optioner[V] {
	return HeadO(vs)
}

func Tail[V any](vs []V) V {
	l := len(vs)

	if l > 0 {
		return vs[l-1]
	} else {
		var zero V
		return zero
	}
}

func TailV[V any](vs ...V) V {
	return Tail(vs)
}

func TailO[V any](vs []V) optioner.Optioner[V] {
	l := len(vs)

	if l > 0 {
		return optioner.Some(vs[l-1])
	} else {
		return optioner.None[V]()
	}
}

func TailVO[V any](vs ...V) optioner.Optioner[V] {
	return TailO(vs)
}
