package optioner

// Some[T] creates an Option[T] with the given value.
func Some[T any](v T) Option[T] {
	var o Option[T]
	o.v = &v
	return o
}

// None[T] creates an Option[T] with no value.
func None[T any]() Option[T] {
	var o Option[T]
	return o
}

// OfP[T] creates a Option[T] that may or may not have a value.
// If *T is nil the returned Option[T] is empty (equivalent to None[T]()),
// If *T is non-nil returned Option[T] will contain a value (equivalent to Some[T](v)).
func OfP[T any](v *T) Option[T] {
	return Option[T]{
		v: v,
	}
}

// OfZ[T] creates a Option[T] that may or may not have a value.
// If T is the zero value, the returned Option[T] is empty (equivalent to None[T]()),
// If T is not a zero, the returned Option[T] will contain a value (equivalent to Some[T](v)).
func OfZ[T comparable](v T) Option[T] {
	var zero T

	if v == zero {
		return None[T]()
	}

	return Some(v)
}

// OfI[T] creates a Option[T] that may or may not have a value.
// If ok is false, the returned Option[T] is empty (equivalent to None[T]()),
// If ok is true, the returned Option[T] will contain the value v (equivalent to Some[T](v)).
func OfI[T any](v T, ok bool) Option[T] {
	if ok {
		return Some(v)
	}

	return None[T]()
}
