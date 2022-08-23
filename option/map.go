package option

func ToFlatMapFunc[A any, B any](fn func(A) B) func(A) Option[B] {
	return func(a A) Option[B] {
		b := fn(a)
		return Some(b)
	}
}

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
