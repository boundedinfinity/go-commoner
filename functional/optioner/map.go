package optioner

func FlatMap[A any, B any](a Option[A], fn func(A) Option[B]) Option[B] {
	if a.Defined() {
		return fn(a.Get())
	}

	return None[B]()
}

func Map[A any, B any](a Option[A], fn func(A) B) Option[B] {
	if a.Defined() {
		b := fn(a.Get())
		return Some(b)
	}

	return None[B]()
}
