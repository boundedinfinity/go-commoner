package option

import "github.com/boundedinfinity/go-commoner/slices"

func FlatMap[A any, B any](as []Option[A], fn func(A) Option[B]) []Option[B] {
	fn2 := func(a Option[A]) Option[B] {
		if a.Defined() {
			return fn(a.Get())
		} else {
			return None[B]()
		}
	}

	return slices.Map(as, fn2)
}
