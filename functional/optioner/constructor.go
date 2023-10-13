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

// OfPtr[T] creates a Option[T] that may or may not have a value.
// If *T is nil the returned Option[T] is empty (equivalent to None[T]()),
// If *T is non-nil returned Option[T] will contain a value (equivalent to Some[T](v)).
func OfPtr[T any](v *T) Option[T] {
	return Option[T]{
		v: v,
	}
}

// OfZero[T] creates a Option[T] that may or may not have a value.
// If T is the zero value, the returned Option[T] is empty (equivalent to None[T]()),
// If T is not a zero, the returned Option[T] will contain a value (equivalent to Some[T](v)).
func OfZero[T comparable](v T) Option[T] {
	var zero T

	if v == zero {
		return None[T]()
	}

	return Some(v)
}

// OfFn[T] creates a Option[T] that may or may not have a value.
// If fn(V) is false, the returned Option[T] is empty (equivalent to None[T]()),
// If fn(V) is true, the returned Option[T] will contain the value v (equivalent to Some[T](v)).
func OfFn[T any](v T, fn func(v T) bool) Option[T] {
	if fn(v) {
		return Some(v)
	}

	return None[T]()
}

// OfOk[T] creates a Option[T] that may or may not have a value.
// If ok is false, the returned Option[T] is empty (equivalent to None[T]()),
// If ok is true, the returned Option[T] will contain the value v (equivalent to Some[T](v)).
func OfOk[T any](v T, ok bool) Option[T] {
	return OfFn(v, func(v T) bool {
		return ok
	})
}

// OfS[[]T] creates a Option[[]T] that may or may not have a value.
//
// If len([]T) > 0  the returned Option[[]T] will contain the value v (equivalent to Some[[]T](v))
//
// Otherwise the returned Option[[]T] is empty, equivalent to None[[]T]().
func OfS[T any](v []T) Option[[]T] {
	return OfFn(v, func(v []T) bool {
		return len(v) > 0
	})
}

// OfM[map[K]V] creates a Option[map[K]V] that may or may not have a value.
//
// If len(map[K]V) > 0  the returned Option[map[K]V] will contain the value v, equivalent to Some[map[K]V](v)
//
// Otherwise the returned Option[map[K]V] is empty, equivalent to None[map[K]V]().
func OfM[K comparable, V any](v map[K]V) Option[map[K]V] {
	return OfFn(v, func(v map[K]V) bool {
		return len(v) > 0
	})
}
