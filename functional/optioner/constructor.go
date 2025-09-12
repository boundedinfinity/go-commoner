package optioner

// Some creates an Option[T] with the given value.
func Some[T any](v T) Option[T] {
	var o Option[T]
	o.v = &v
	return o
}

// None creates an Option[T] with no value.
func None[T any]() Option[T] {
	var o Option[T]
	return o
}

// OfPtr creates a Option[T] that may or may not have a value.
//
// If *T is nil None[T] is returned.
//
// If *T is not nil Some[T] is returned.
func OfPtr[T any](v *T) Option[T] {
	return Option[T]{
		v: v,
	}
}

// OfZero creates a Option[T] that may or may not have a value.
//
// If T is equivalent the zero value of T None[T] is returned.
//
// If T is not equivalent the zero value of T Some[T] is returned.
func OfZero[T comparable](v T) Option[T] {
	var zero T

	if v == zero {
		return None[T]()
	}

	return Some(v)
}

// OfFn creates a Option[T] that may or may not have a value.
//
// If fn(V) is false, None[T] is returned.
//
// If fn(V) is true, Some[T] is returned.
func OfFn[T any](v T, fn func(v T) bool) Option[T] {
	if fn(v) {
		return Some(v)
	}

	return None[T]()
}

// OfOk creates a Option[T] that may or may not have a value.
//
// If ok is false, None[T] is returned.
//
// If ok is true, Some[T] is returned.
func OfOk[T any](v T, ok bool) Option[T] {
	return OfFn(v, func(v T) bool {
		return ok
	})
}

// OfSlice creates a Option[[]T] that may or may not have a value.
//
// If len([]T) > 0, Some[[]T] is returned.
//
// if len([]T) == 0 None[[]T] is return.
func OfSlice[T any](v []T) Option[[]T] {
	return OfFn(v, func(v []T) bool {
		return len(v) > 0
	})
}

// OfMap creates a Option[map[K]V] that may or may not have a value.
//
// If len(map[K]V) > 0, Some[map[K]V] is returned.
//
// If len(map[K]V) == 0, None[map[K]V] is returned.
func OfMap[K comparable, V any](v map[K]V) Option[map[K]V] {
	return OfFn(
		v,
		func(v map[K]V) bool { return len(v) > 0 },
	)
}

// OfOptions creates a Option[[]T] that may or may not have a value.
func OfOptions[T any](elems ...Option[T]) Option[[]T] {
	var realItems []T

	for _, elem := range elems {
		if elem.Defined() {
			realItems = append(realItems, elem.Get())
		}
	}

	return OfSlice(realItems)
}
