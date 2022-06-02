package slices

type Numeric interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Sum[V Numeric](vs []V) V {
	fn := func(a V, v V) V {
		return a + v
	}

	return Reduce(0, vs, fn)
}

func Mean[V Numeric](vs []V) V {
	l := len(vs)

	if l == 0 {
		return 0
	}

	sum := Sum(vs)
	mean := sum / V(l)

	return mean
}

func Min[V Numeric](vs []V) V {
	fn := func(a V, v V) V {
		if v < a {
			return v
		} else {
			return a
		}
	}

	return Reduce(Head(vs...), vs, fn)
}

func Max[V Numeric](vs []V) V {
	fn := func(a V, v V) V {
		if v < a {
			return v
		} else {
			return a
		}
	}

	return Reduce(Head(vs...), vs, fn)
}
